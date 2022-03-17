package storage

import (
	"github.com/daison12006013/gorvel/pkg/facade/path"
	"io"
	"mime/multipart"
	"os"
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

func (s *LocalStorage) Put(path string, file multipart.FileHeader) error {
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
