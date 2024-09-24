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
	if err := os.Mkdir(dir, 0744); err != nil && !errors.Is(err, os.ErrExist) {
		return nil, err
	}
	return &FilesystemRepo{dir: dir}, nil
}

func (r *FilesystemRepo) CreateFile(userID int, reader io.Reader) (uuid.UUID, error) {
	var id uuid.UUID
	for id = uuid.New(); r.isExist(r.filepath(userID, id)); id = uuid.New() {
	}
	err := os.MkdirAll(r.userDir(userID), 0744)
	if err != nil {
		return uuid.Nil, err
	}
	f, err := os.Create(r.filepath(userID, id))
	if err != nil {
		return uuid.Nil, nil
	}
	defer f.Close()
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

func (r *FilesystemRepo) userDir(userID int) string {
	return path.Join(r.dir, strconv.Itoa(userID))
}

func (r *FilesystemRepo) filepath(userID int, id uuid.UUID) string {
	return path.Join(r.userDir(userID), id.String())
}
