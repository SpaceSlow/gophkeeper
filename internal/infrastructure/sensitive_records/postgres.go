package sensitive_records

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
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

func (r *PostgresRepo) CreateSensitiveRecord(record *sensitive_records.SensitiveRecord) (*sensitive_records.SensitiveRecord, error) {
	row := r.pool.QueryRow(
		r.ctx,
		`INSERT INTO sensitive_records (user_id, sensitive_record_type_id, metadata) VALUES ($1, $2, $3) RETURNING id`,
		record.UserID(), record.TypeID(), record.Metadata(),
	)
	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	newRecord, err := sensitive_records.NewSensitiveRecord(id, record.UserID(), record.TypeID(), record.Metadata())
	if err != nil {
		return nil, err
	}
	return newRecord, nil
}

func (r *PostgresRepo) ListSensitiveRecordTypes() ([]sensitive_records.SensitiveRecordType, error) {
	rows, err := r.pool.Query(
		r.ctx,
		"SELECT id, name FROM sensitive_record_types",
	)

	if err != nil {
		return nil, err
	}

	types := make([]sensitive_records.SensitiveRecordType, 0)
	for rows.Next() {
		var (
			id   int
			name string
		)

		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		t := sensitive_records.NewSensitiveRecordType(id, name)
		types = append(types, *t)
	}
	return types, nil
}
func (r *PostgresRepo) Close() {
	r.pool.Close()
}
