package sensitive_records

import "errors"

var ErrInvalidSensitiveRecordType = errors.New("incorrect sensitive record type")

type SensitiveRecordTypeID int

const (
	_ SensitiveRecordTypeID = iota
	Text
	BinaryFile
	Credential
	PaymentCard
)

func NewSensitiveRecordTypeID(srType string) (SensitiveRecordTypeID, error) {
	switch srType {
	case "text":
		return Text, nil
	case "binary-file":
		return BinaryFile, nil
	case "credential":
		return Credential, nil
	case "payment-card":
		return PaymentCard, nil
	default:
		return 0, ErrInvalidSensitiveRecordType
	}
}

type SensitiveRecordType struct {
	id   int
	name string
}

func NewSensitiveRecordType(id int, name string) *SensitiveRecordType {
	return &SensitiveRecordType{
		id:   id,
		name: name,
	}
}

func (t *SensitiveRecordType) Id() int {
	return t.id
}

func (t *SensitiveRecordType) Name() string {
	return t.name
}
