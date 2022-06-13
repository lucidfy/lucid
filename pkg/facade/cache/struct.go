package cache

import "github.com/lucidfy/lucid/pkg/errors"

type CacheContract interface {
	Put(name string, value interface{}) (bool, *errors.AppError)
	Get(name string) (interface{}, *errors.AppError)
	GetAs(name string, m interface{})
	Forget(name string) (interface{}, *errors.AppError)
}

func Store(key string) CacheContract {
	switch key {
	case "file":
		return File()
	}
	return nil
}
