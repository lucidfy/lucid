package storage

import "mime/multipart"

type Storage interface {
	Get(path string) (multipart.File, error)
	Put(path string, file *multipart.FileHeader) error
	Delete(path string) error

	Exists(path string) bool
	Missing(path string) bool
	Size(path string) int64

	Path(path string) (string, bool)
}
