package sensitive_records

type SensitiveRecord struct {
	id       int
	userID   int
	dType    string
	metadata string
}

func NewSensitiveRecord(id, userID int, dType string, metadata string) (*SensitiveRecord, error) {
	return &SensitiveRecord{
		id:       id,
		userID:   userID,
		dType:    dType,
		metadata: metadata,
	}, nil
}

func CreateSensitiveRecord(userID int, dType string, metadata string) (*SensitiveRecord, error) {
	return NewSensitiveRecord(0, userID, dType, metadata)
}

func (r *SensitiveRecord) Id() int {
	return r.id
}

func (r *SensitiveRecord) UserID() int {
	return r.userID
}

func (r *SensitiveRecord) Type() string {
	return r.dType
}

func (r *SensitiveRecord) Metadata() string {
	return r.metadata
}
