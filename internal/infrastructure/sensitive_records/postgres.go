package sensitive_records

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *PostgresRepo) UploadSensitiveRecord() (bool, error) {
	return false, errors.New("not implemented")
}

func (r *PostgresRepo) Close() {
	r.pool.Close()
}
