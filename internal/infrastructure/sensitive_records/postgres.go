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
		`INSERT INTO sensitive_records (user_id, type, metadata) VALUES ($1, $2, $3) RETURNING id`,
		record.UserID(), record.Type(), record.Metadata(),
	)
	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	newRecord, err := sensitive_records.NewSensitiveRecord(id, record.UserID(), record.Type(), record.Metadata())
	if err != nil {
		return nil, err
	}
	return newRecord, nil
}

func (r *PostgresRepo) ListSensitiveRecords(userID int) ([]sensitive_records.SensitiveRecord, error) {
	rows, err := r.pool.Query(
		r.ctx,
		`SELECT id, type, metadata FROM sensitive_records WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	var (
		id       int
		dType    string
		metadata string
	)
	records := make([]sensitive_records.SensitiveRecord, 0)
	for rows.Next() {
		err := rows.Scan(&id, &dType, &metadata)
		if err != nil {
			return nil, err
		}
		record, err := sensitive_records.NewSensitiveRecord(id, userID, dType, metadata)
		if err != nil {
			return nil, err
		}
		records = append(records, *record)
	}

	return records, nil
}

func (r *PostgresRepo) Close() {
	r.pool.Close()
}
