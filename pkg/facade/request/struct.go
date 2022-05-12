package request

import (
	"mime/multipart"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type RequestContract interface {
	All() interface{}
	Get(k string) interface{}
	GetFirst(k string, dfault interface{}) interface{}
	Input(k string, dfault interface{}) interface{}
	HasContentType(substr string) bool
	HasAccept(substr string) bool
	IsForm() bool
	IsJson() bool
	IsMultipart() bool
	WantsJson() bool
	Validator(setOfRules *must.SetOfRules) *errors.AppError
	GetIp() string
	GetUserAgent() string
	GetFileByName(name string) (*multipart.FileHeader, error)
	GetFiles() (map[string][]*multipart.FileHeader, error)
}
