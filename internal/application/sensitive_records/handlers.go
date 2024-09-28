package sensitive_records

import "github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"

type Repository interface {
	CreateSensitiveRecord(sensitiveRecord *sensitive_records.SensitiveRecord) (*sensitive_records.SensitiveRecord, error)
	ListSensitiveRecords(userID int) ([]sensitive_records.SensitiveRecord, error)
	Close()
}

type SensitiveRecordHandlers struct {
	repo Repository
}

func SetupHandlers(repo Repository) SensitiveRecordHandlers {
	return SensitiveRecordHandlers{
		repo: repo,
	}
}
