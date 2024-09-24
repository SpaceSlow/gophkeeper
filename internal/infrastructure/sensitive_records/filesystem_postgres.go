package sensitive_records

import "context"

type FilesystemPostgresRepo struct {
	*FilesystemRepo
	*PostgresRepo
}

func NewFilesystemPostgresRepo(ctx context.Context, dsn, dir string) (*FilesystemPostgresRepo, error) {
	filesystemRepo, err := NewFilesystemRepo(dir)
	if err != nil {
		return nil, err
	}
	postgresRepo, err := NewPostgresRepo(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &FilesystemPostgresRepo{FilesystemRepo: filesystemRepo, PostgresRepo: postgresRepo}, nil
}
