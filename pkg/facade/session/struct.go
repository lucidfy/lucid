package session

import (
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

func Driver(key string, sessionKey string) SessionContract {
	switch key {
	case "file":
		return File(sessionKey)
	}
	return nil
}
