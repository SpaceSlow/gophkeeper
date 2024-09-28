package sensitive_records

import "io"

type SensitiveRecordData struct {
	sensitiveRecordID int
	data              io.Reader
}

func NewSensitiveRecordData(sensitiveRecordID int, data io.Reader) (*SensitiveRecordData, error) {
	return &SensitiveRecordData{
		sensitiveRecordID: sensitiveRecordID,
		data:              data,
	}, nil
}

func (d *SensitiveRecordData) SensitiveRecordID() int {
	return d.sensitiveRecordID
}

func (d *SensitiveRecordData) DataAsBytes() ([]byte, error) {
	return io.ReadAll(d.data)
}
