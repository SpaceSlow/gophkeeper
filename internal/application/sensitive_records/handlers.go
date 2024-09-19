package sensitive_records

type Repository interface {
	UploadSensitiveRecord() (bool, error)
	Close()
}

type SensitiveRecordHandlers struct {
	repo     Repository
	strategy SensitiveRecordStrategy
}

func SetupHandlers(repo Repository) SensitiveRecordHandlers {
	return SensitiveRecordHandlers{
		repo: repo,
	}
}
