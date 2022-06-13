package cookie

import "github.com/lucidfy/lucid/pkg/errors"

type CookieContract interface {
	CreateSessionCookie() interface{}
	Set(name string, value interface{}) (bool, *errors.AppError)
	Get(name string) (interface{}, *errors.AppError)
	Expire(string)
}
