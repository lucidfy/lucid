package storage

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/path"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		basePath: path.Load().StoragePath(""),
	}
}

func (s *LocalStorage) Get(path string) (multipart.File, *errors.AppError) {
	f, err := os.Open(s.basePath + "/" + path)
	return f, errors.InternalServerError("os.Open() error", err)
}

func (s *LocalStorage) Put(path string, file *multipart.FileHeader) *errors.AppError {
	src, err := file.Open()
	if err != nil {
		return errors.InternalServerError("file.Open() error", err)
	}

	defer src.Close()

	out, err := os.Create(s.basePath + "/" + path)
	if err != nil {
		return errors.InternalServerError("os.Create() error", err)
	}

	defer out.Close()

	_, err = io.Copy(out, src)
	return errors.InternalServerError("os.Copy() error", err)
}

// Exists check if file exists
func (s *LocalStorage) Exists(path string) bool {
	_, err := os.Stat(s.basePath + "/" + path)
	return !os.IsNotExist(err)
}

// Missing check if file exists
func (s *LocalStorage) Missing(path string) bool {
	_, err := os.Stat(s.basePath + "/" + path)
	return os.IsNotExist(err)
}

// Size get file size
func (s *LocalStorage) Size(path string) int64 {
	fileInfo, err := os.Stat(s.basePath + "/" + path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

// Delete file
func (s *LocalStorage) Delete(path string) *errors.AppError {
	err := os.Remove(s.basePath + "/" + path)
	return errors.InternalServerError("os.Remove() error", err)
}

// Path get file path
func (s *LocalStorage) Path(path string) (string, bool) {
	if s.Missing(path) {
		return "", false
	}
	return s.basePath + "/" + path, true
}
