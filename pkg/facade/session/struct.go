package session

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/errors"
)

type SessionContract interface {
	Put(name string, value interface{}) (bool, *errors.AppError)
	Get(name string) (interface{}, *errors.AppError)
	Flush(name string) (interface{}, *errors.AppError)
	PutFlash(name string, value interface{})
	GetFlash(name string) interface{}
	PutFlashMap(name string, value interface{})
	GetFlashMap(name string) *map[string]interface{}
}

func Driver(key string, w http.ResponseWriter, r *http.Request) SessionContract {
	switch key {
	case "file":
		return File(w, r)
	}
	return nil
}
