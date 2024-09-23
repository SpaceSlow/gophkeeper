package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
)

type PostgresRepo struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, dsn string) (*PostgresRepo, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create a connection pool: %w", err)
	}
	return &PostgresRepo{
		pool: pool,
		ctx:  ctx,
	}, nil
}

func (r *PostgresRepo) ExistUsername(username string) (bool, error) {
	row := r.pool.QueryRow(
		r.ctx,
		`SELECT EXISTS(SELECT id FROM users WHERE username=$1)`,
		username,
	)
	var existUsername bool
	if err := row.Scan(&existUsername); err != nil {
		return false, fmt.Errorf("failed to check existing username: %w", err)
	}
	return existUsername, nil
}

func (r *PostgresRepo) RegisterUser(username, passwordHash string) error {
	_, err := r.pool.Exec(
		r.ctx,
		`INSERT INTO users (username, password_hash) VALUES ($1, $2)`,
		username, passwordHash,
	)
	return err
}
func (r *PostgresRepo) FetchPasswordHash(username string) (string, error) {
	row := r.pool.QueryRow(
		r.ctx,
		"SELECT password_hash FROM users WHERE username=$1",
		username,
	)

	var hash string
	err := row.Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", users.NewNoUserError(username)
	} else if err != nil {
		return "", err
	}
	return hash, nil
}

func (r *PostgresRepo) FetchUserID(username string) (int, error) {
	row := r.pool.QueryRow(
		r.ctx,
		"SELECT id FROM users WHERE username=$1",
		username,
	)

	var userID int
	err := row.Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		return -1, users.NewNoUserError(username)
	} else if err != nil {
		return -1, err
	}
	return userID, nil
}

func (r *PostgresRepo) Close() {
	r.pool.Close()
}
