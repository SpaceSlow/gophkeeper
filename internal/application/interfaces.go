package application

import "github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"

type SensitiveRecordRepository interface {
	ListSensitiveRecordTypes() ([]sensitive_records.SensitiveRecordType, error)
	UploadSensitiveRecord() (bool, error)
	Close()
}

type UserRepository interface {
	ExistUsername(username string) (bool, error)
	RegisterUser(username, passwordHash string) error
	FetchPasswordHash(username string) (string, error)
	FetchUserID(username string) (int, error)
	Close()
}
