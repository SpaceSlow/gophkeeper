package application

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/SpaceSlow/gophkeeper/internal"
	"github.com/SpaceSlow/gophkeeper/internal/infrastructure/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/infrastructure/users"
)

type Server struct {
	ctx    context.Context
	config *internal.ServerConfig

	srv *http.Server
}

func NewServer() (*Server, error) {
	var srv Server
	srv.ctx = context.Background()
	srv.config = internal.LoadServerConfig()
	return &srv, nil
}

func (s *Server) Run() error {
	rootCtx, cancelCtx := signal.NotifyContext(s.ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancelCtx()

	g, ctx := errgroup.WithContext(rootCtx)
	s.ctx = ctx

	context.AfterFunc(ctx, func() {
		timeoutCtx, cancelCtx := context.WithTimeout(context.Background(), s.config.TimeoutShutdown)
		defer cancelCtx()

		<-timeoutCtx.Done()
		slog.Error("failed to gracefully shutdown the service")
	})

	if err := internal.RunMigrations(s.config.DSN); err != nil {
		return fmt.Errorf("failed to run DB migrations: %w", err)
	}

	userRepo, err := users.NewPostgresRepo(ctx, s.config.DSN)
	if err != nil {
		return fmt.Errorf("failed to initialize a user repo: %w", err)
	}
	defer userRepo.Close()

	g.Go(func() error {
		defer slog.Info("closed user repo")
		<-ctx.Done()
		userRepo.Close()
		return nil
	})

	sensitiveRecordRepo, err := sensitive_records.NewPostgresRepo(ctx, s.config.DSN)
	if err != nil {
		return fmt.Errorf("failed to initialize a sensitive record repo: %w", err)
	}
	defer sensitiveRecordRepo.Close()

	g.Go(func() error {
		defer slog.Info("closed sensitive record repo")
		<-ctx.Done()
		sensitiveRecordRepo.Close()
		return nil
	})

	g.Go(func() (err error) {
		defer func() {
			errRec := recover()
			if errRec != nil {
				err = fmt.Errorf("a panic occurred: %v", errRec)
			}
		}()

		s.srv = &http.Server{
			Addr:         s.config.NetAddress.String(),
			Handler:      SetupRouter(userRepo, sensitiveRecordRepo), // TODO: fix
			TLSConfig:    s.tlsConfig(),
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		return s.srv.ListenAndServeTLS(s.config.CertificatePath, s.config.PrivateKeyPath)
	})

	g.Go(func() error {
		defer slog.Info("server has been shutdown")

		<-ctx.Done()

		shutdownTimeoutCtx, cancelShutdownTimeoutCtx := context.WithTimeout(context.Background(), s.config.TimeoutShutdown)
		defer cancelShutdownTimeoutCtx()
		return s.srv.Shutdown(shutdownTimeoutCtx)
	})

	if err := g.Wait(); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func (s *Server) tlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
}
