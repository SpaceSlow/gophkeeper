package sensitive_records

import "fmt"

type ExistSensitiveRecordDataError struct {
	id int
}

func NewExistSensitiveRecordDataError(id int) error {
	return &ExistSensitiveRecordDataError{id: id}
}

func (e *ExistSensitiveRecordDataError) Error() string {
	return fmt.Sprintf("already exist data of sensitive record with id=%d", e.id)
}

type NotExistSensitiveRecordDataError struct {
	id int
}

func NewNotExistSensitiveRecordDataError(id int) error {
	return &NotExistSensitiveRecordDataError{id: id}
}

func (e *NotExistSensitiveRecordDataError) Error() string {
	return fmt.Sprintf("not exist data of sensitive record with id=%d", e.id)
}
