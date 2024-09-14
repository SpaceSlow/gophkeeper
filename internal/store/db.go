package store

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context, dsn string) (*DB, error) {
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("failed to run DB migrations: %w", err)
	}
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create a connection pool: %w", err)
	}
	return &DB{
		pool: pool,
	}, nil
}

func (db *DB) CheckUsername(ctx context.Context, username string) (bool, error) {
	row := db.pool.QueryRow(
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

func (db *DB) RegisterUser(ctx context.Context, username, passwordHash string) error {
	_, err := db.pool.Exec(
		ctx,
		`INSERT INTO users (username, password_hash) VALUES ($1, $2)`,
		username, passwordHash,
	)
	return err
}

func (db *DB) FetchPasswordHash(ctx context.Context, username string) (string, error) {
	row := db.pool.QueryRow(
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

func (db *DB) FetchUserID(ctx context.Context, username string) (int, error) {
	row := db.pool.QueryRow(
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

func (db *DB) Close() {
	db.pool.Close()
}
