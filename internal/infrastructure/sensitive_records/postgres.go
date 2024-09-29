package sensitive_records

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *PostgresRepo) DeleteSensitiveRecord(id int) error {
	tx, err := r.pool.BeginTx(r.ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(r.ctx, `DELETE FROM sensitive_datas WHERE sensitive_record_id=$1`, id)
	if err != nil {
		tx.Rollback(r.ctx)
		return err
	}
	_, err = tx.Exec(r.ctx, `DELETE FROM sensitive_records WHERE id=$1`, id)
	if err != nil {
		tx.Rollback(r.ctx)
		return err
	}
	return tx.Commit(r.ctx)
}

func (r *PostgresRepo) IsSensitiveRecordOwner(id, userID int) (bool, error) {
	row := r.pool.QueryRow(
		r.ctx,
		`SELECT EXISTS(SELECT 1 FROM sensitive_records WHERE id=$1 AND user_id=$2)`,
		id, userID,
	)

	var isOwner bool
	err := row.Scan(&isOwner)
	return isOwner, err
}

func (r *PostgresRepo) CreateSensitiveRecordData(data *sensitive_records.SensitiveRecordData) error {
	_, err := r.pool.Exec(
		r.ctx,
		`INSERT INTO sensitive_datas(sensitive_record_id, data) VALUES ($1, $2)`,
		data.SensitiveRecordID(), data.Data(),
	)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		return sensitive_records.NewExistSensitiveRecordDataError(data.SensitiveRecordID())
	}
	return err
}

func (r *PostgresRepo) FetchSensitiveRecordData(id int) (*sensitive_records.SensitiveRecordData, error) {
	row := r.pool.QueryRow(
		r.ctx,
		`SELECT data FROM sensitive_datas WHERE sensitive_record_id=$1`,
		id,
	)
	var data []byte
	if err := row.Scan(&data); errors.Is(err, pgx.ErrNoRows) {
		return nil, sensitive_records.NewNotExistSensitiveRecordDataError(id)
	} else if err != nil {
		return nil, err
	}
	sensitiveRecordData, err := sensitive_records.NewSensitiveRecordData(id, data)
	return sensitiveRecordData, err
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
