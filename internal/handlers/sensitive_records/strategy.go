package sensitive_records

import (
	"fmt"
	"net/http"
)

type SensitiveRecordStrategy interface {
	Upload(req *http.Request) (string, error)
}

func NewSensitiveRecordStrategy(recordType string) (SensitiveRecordStrategy, error) {
	switch recordType {
	case "text-file":
		return &TextFileStrategy{}, nil
	case "binary-file":
		return &BinaryFileStrategy{}, nil
	case "credentials":
		return &LoginPasswordStrategy{}, nil
	case "payment-card":
		return &LoginPasswordStrategy{}, nil
	default:
		return nil, fmt.Errorf("unknown sensitive record type: %s", recordType)
	}
}
