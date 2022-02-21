package request

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/daison12006013/gorvel/pkg/facade/urls"
)

type ParsedRequest struct {
	HttpRequest *http.Request
}

func Parse(r *http.Request) ParsedRequest {
	t := ParsedRequest{
		HttpRequest: r,
	}
	return t
}

// Url  ----------------------------------------------------

func (t ParsedRequest) CurrentUrl() string {
	h := t.HttpRequest.URL.Host
	if len(h) > 0 {
		return h
	}
	// if ever we can't resolve the Host from net/http, we can still base
	// from our env config
	return urls.BaseUrl(nil)
}

func (t ParsedRequest) FullUrl() string {
	return t.CurrentUrl() + t.HttpRequest.URL.RequestURI()
}

func (t ParsedRequest) PreviousUrl() string {
	return t.HttpRequest.Referer()
}

// Request  -------------------------------------------------

// This returns all avaiable queries from
func (t ParsedRequest) All() url.Values {
	params := t.HttpRequest.URL.Query()
	return params
}

// This returns the specific value from the provided key
func (t ParsedRequest) Get(k string) []string {
	val, ok := t.All()[k]
	if ok {
		return val
	}
	return []string{}
}

func (t ParsedRequest) GetFirst(k string, dfault *string) *string {
	val := t.Get(k)
	if len(val) > 0 {
		return &val[0]
	}
	return dfault
}

// Proxy method to GetFirst(...)
func (t ParsedRequest) Input(k string, dfault *string) *string {
	return t.GetFirst(k, dfault)
}

// Check if the string exists in the content type
func (t ParsedRequest) HasContentType(substr string) bool {
	contentType := t.HttpRequest.Header.Get("Content-Type")
	return strings.Contains(contentType, substr)
}

func (t ParsedRequest) HasAccept(substr string) bool {
	accept := t.HttpRequest.Header.Get("Accept")
	return strings.Contains(accept, substr)
}

// Check if the request is form
func (t ParsedRequest) IsForm() bool {
	return t.HasContentType("application/x-www-form-urlencoded")
}

// Check if the request is json
func (t ParsedRequest) IsJson() bool {
	return t.HasContentType("json")
}

// Check if the request is multipart
func (t ParsedRequest) IsMultipart() bool {
	return t.HasContentType("multipart")
}

func (t ParsedRequest) WantsJson() bool {
	return t.HasAccept("json")
}
