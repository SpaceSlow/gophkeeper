package sensitive_records

import (
	"errors"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/google/uuid"
)

type FilesystemRepo struct {
	dir string
}

func NewFilesystemRepo(dir string) (*FilesystemRepo, error) {
	if err := os.Mkdir(dir, 644); err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}
	return &FilesystemRepo{dir: dir}, nil
}

func (r *FilesystemRepo) CreateBinaryFile(userID int, reader io.Reader) (uuid.UUID, error) {
	var id uuid.UUID
	for id = uuid.New(); r.isExist(r.path(userID, id)); id = uuid.New() {
	}
	f, err := os.Open(r.path(userID, id))
	defer f.Close()

	if err != nil {
		return uuid.Nil, nil
	}
	_, err = io.Copy(f, reader)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *FilesystemRepo) isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (r *FilesystemRepo) path(userID int, id uuid.UUID) string {
	return path.Join(r.dir, strconv.Itoa(userID), id.String())
}
