package application

import (
	"io"

	"github.com/google/uuid"

	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
)

type SensitiveRecordRepository interface {
	ListSensitiveRecordTypes() ([]sensitive_records.SensitiveRecordType, error)
	CreateSensitiveRecord(sensitiveRecord *sensitive_records.SensitiveRecord) (*sensitive_records.SensitiveRecord, error)
	CreateFile(userID int, reader io.Reader) (uuid.UUID, error)
	Close()
}

type UserRepository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchUser(username string) (*users.User, error)
	ExistUser(userID int) (bool, error)
	Close()
}
