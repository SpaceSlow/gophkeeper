package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/SpaceSlow/gophkeeper/internal/application"
	"github.com/SpaceSlow/gophkeeper/internal/infrastructure/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/infrastructure/users"
)

func RunServer() error {
	rootCtx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancelCtx()

	cfg := LoadServerConfig()
	g, ctx := errgroup.WithContext(rootCtx)
	context.AfterFunc(ctx, func() {
		timeoutCtx, cancelCtx := context.WithTimeout(context.Background(), cfg.TimeoutShutdown)
		defer cancelCtx()

		<-timeoutCtx.Done()
		slog.Error("failed to gracefully shutdown the service")
	})

	// run migrations
	if err := RunMigrations(cfg.DSN); err != nil {
		return fmt.Errorf("failed to run DB migrations: %w", err)
	}

	// start sensitive record repository
	userRepo, err := users.NewPostgresRepo(ctx, cfg.DSN)
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

	// start sensitive record repository
	sensitiveRecordRepo, err := sensitive_records.NewPostgresRepo(ctx, cfg.DSN)
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

	// setup and start https server
	httpServer := application.SetupHTTPServer(userRepo, sensitiveRecordRepo, cfg)
	srv := &http.Server{
		Addr:         cfg.NetAddress.String(),
		Handler:      httpServer,
		TLSConfig:    tlsConfig(),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	g.Go(func() (err error) {
		defer func() {
			errRec := recover()
			if errRec != nil {
				err = fmt.Errorf("a panic occurred: %v", errRec)
			}
		}()
		return srv.ListenAndServeTLS(cfg.CertificatePath, cfg.PrivateKeyPath)
	})

	// server closed
	g.Go(func() error {
		defer slog.Info("server has been shutdown")

		<-ctx.Done()

		shutdownTimeoutCtx, cancelShutdownTimeoutCtx := context.WithTimeout(context.Background(), cfg.TimeoutShutdown)
		defer cancelShutdownTimeoutCtx()
		return srv.Shutdown(shutdownTimeoutCtx)
	})

	if err := g.Wait(); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

func tlsConfig() *tls.Config {
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
