package request

import (
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/rules/must"
	"mime/multipart"
)

type RequestContract interface {
	CurrentUrl() string
	FullUrl() string
	PreviousUrl() string
	RedirectPrevious()
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
}
