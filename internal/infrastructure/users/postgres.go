package users

import (
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, dsn string) (*PostgresRepo, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("failed to run DB migrations: %w", err)
	}
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create a connection pool: %w", err)
	}
	return &PostgresRepo{
		pool: pool,
	}, nil
}

func (r *PostgresRepo) ExistUsername(ctx context.Context, username string) (bool, error) {
	row := r.pool.QueryRow(
		ctx,
		`SELECT EXISTS(SELECT id FROM users WHERE username=$1)`,
		username,
	)
	var existUsername bool
	if err := row.Scan(&existUsername); err != nil {
		return false, fmt.Errorf("failed to check existing username: %w", err)
	}
	return existUsername, nil
}

func (r *PostgresRepo) RegisterUser(ctx context.Context, username, passwordHash string) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO users (username, password_hash) VALUES ($1, $2)`,
		username, passwordHash,
	)
	return err
}
func (r *PostgresRepo) FetchPasswordHash(ctx context.Context, username string) (string, error) {
	row := r.pool.QueryRow(
		ctx,
		"SELECT password_hash FROM users WHERE username=$1",
		username,
	)

	var hash string
	err := row.Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", &NoUserError{Username: username}
	} else if err != nil {
		return "", err
	}
	return hash, nil
}

func (r *PostgresRepo) FetchUserID(ctx context.Context, username string) (int, error) {
	row := r.pool.QueryRow(
		ctx,
		"SELECT id FROM users WHERE username=$1",
		username,
	)

	var userID int
	err := row.Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return -1, &NoUserError{Username: username}
	} else if err != nil {
		return -1, err
	}
	return userID, nil
}

func (r *PostgresRepo) Close() {
	r.pool.Close()
}

//go:embed migrations/*.sql
var migrationsDir embed.FS

func runMigrations(dsn string) error {
	d, err := iofs.New(migrationsDir, "migrations")
	if err != nil {
		return fmt.Errorf("failed to return an iofs driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return fmt.Errorf("failed to get a new migrate instance: %w", err)
	}
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("failed to apply migrations to the DB: %w", err)
		}
	}
	return nil
}
