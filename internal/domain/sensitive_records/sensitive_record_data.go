package sensitive_records

type SensitiveRecordData struct {
	sensitiveRecordID int
	data              []byte
}

func NewSensitiveRecordData(sensitiveRecordID int, data []byte) (*SensitiveRecordData, error) {
	return &SensitiveRecordData{
		sensitiveRecordID: sensitiveRecordID,
		data:              data,
	}, nil
}

func (d *SensitiveRecordData) SensitiveRecordID() int {
	return d.sensitiveRecordID
}

func (d *SensitiveRecordData) Data() []byte {
	return d.data
}
