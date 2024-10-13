package application

import (
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/domain/users"
)

type SensitiveRecordRepository interface {
	CreateSensitiveRecord(sensitiveRecord *sensitive_records.SensitiveRecord) (*sensitive_records.SensitiveRecord, error)
	DeleteSensitiveRecord(id int) error
	CreateSensitiveRecordData(data *sensitive_records.SensitiveRecordData) error
	FetchSensitiveRecordData(id int) (*sensitive_records.SensitiveRecordData, error)
	ListSensitiveRecords(userID int) ([]sensitive_records.SensitiveRecord, error)
	IsSensitiveRecordOwner(id, userID int) (bool, error)
	Close()
}

type UserRepository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchUser(username string) (*users.User, error)
	ExistUser(userID int) (bool, error)
	Close()
}
