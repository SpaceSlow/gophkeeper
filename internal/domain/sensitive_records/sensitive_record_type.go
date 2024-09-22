package sensitive_records

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
