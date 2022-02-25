package request

import "net/url"

type Request interface {
	CurrentUrl() string
	FullUrl() string
	PreviousUrl() string
	RedirectPrevious()
	SetFlash(name string, value string)
	GetFlash(name string) *string
	All() url.Values
	Get(k string) []string
	GetFirst(k string, dfault *string) *string
	Input(k string, dfault string) string
	HasContentType(substr string) bool
	HasAccept(substr string) bool
	IsForm() bool
	IsJson() bool
	IsMultipart() bool
	WantsJson() bool
}
