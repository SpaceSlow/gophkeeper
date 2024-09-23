package sensitive_records

type SensitiveRecord struct {
	id                    int
	userID                int
	sensitiveRecordTypeID int
	metadata              string
}

func NewSensitiveRecord(id, userID, sensitiveRecordTypeID int, metadata string) (*SensitiveRecord, error) {
	return &SensitiveRecord{
		id:                    id,
		userID:                userID,
		sensitiveRecordTypeID: sensitiveRecordTypeID,
		metadata:              metadata,
	}, nil
}

func CreateSensitiveRecord(userID, sensitiveRecordTypeID int, metadata string) (*SensitiveRecord, error) {
	return NewSensitiveRecord(0, userID, sensitiveRecordTypeID, metadata)
}

func (r *SensitiveRecord) Id() int {
	return r.id
}

func (r *SensitiveRecord) UserID() int {
	return r.userID
}

func (r *SensitiveRecord) TypeID() int {
	return r.sensitiveRecordTypeID
}

func (r *SensitiveRecord) Metadata() string {
	return r.metadata
}
