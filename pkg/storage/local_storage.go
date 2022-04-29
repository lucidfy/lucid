package storage

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/daison12006013/lucid/pkg/facade/path"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		basePath: path.Load().StoragePath(""),
	}
}

func (s *LocalStorage) Get(path string) (multipart.File, error) {
	return os.Open(s.basePath + "/" + path)
}

func (s *LocalStorage) Put(path string, file *multipart.FileHeader) error {

	src, err := file.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(s.basePath + "/" + path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
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
func (s *LocalStorage) Delete(path string) error {
	return os.Remove(s.basePath + "/" + path)
}

// Path get file path
func (s *LocalStorage) Path(path string) (string, bool) {
	if s.Missing(path) {
		return "", false
	}
	return s.basePath + "/" + path, true
}
