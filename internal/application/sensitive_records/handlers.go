package sensitive_records

import (
	"fmt"

	"github.com/SpaceSlow/gophkeeper/internal/application/sensitive_records/strategy"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
)

type Repository interface {
	ListSensitiveRecordTypes() ([]sensitive_records.SensitiveRecordType, error)
	UploadSensitiveRecord() (bool, error)
	Close()
}

type SensitiveRecordHandlers struct {
	repo     Repository
	strategy strategy.SensitiveRecordStrategy
}

func SetupHandlers(repo Repository) SensitiveRecordHandlers {
	return SensitiveRecordHandlers{
		repo: repo,
	}
}

func (h *SensitiveRecordHandlers) setStrategy(recordType string) error {
	switch recordType {
	case "text":
		h.strategy = &strategy.TextStrategy{}
	case "binary-file":
		h.strategy = &strategy.BinaryFileStrategy{}
	case "credentials":
		h.strategy = &strategy.CredentialStrategy{}
	case "payment-card":
		h.strategy = &strategy.PaymentCardStrategy{}
	default:
		return fmt.Errorf("unknown sensitive record type: %s", recordType)
	}
	return nil
}
