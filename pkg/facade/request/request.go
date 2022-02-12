package request

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type reqStruct struct {
	httpRequest *http.Request
}

func Parse(r *http.Request) reqStruct {
	t := reqStruct{
		httpRequest: r,
	}
	return t
}

// This returns all avaiable queries from
func (t reqStruct) All() map[string]string {
	params := mux.Vars(t.httpRequest)
	return params
}

// This returns the specific value from the provided key
func (t reqStruct) Get(k string) *string {
	val, ok := t.All()[k]
	if ok {
		return &val
	}
	return nil
}

// Proxy method similar to Get(...)
func (t reqStruct) Input(k string) *string {
	return t.Get(k)
}

// Check if the string exists in the content type
func (t reqStruct) HasContentType(substr string) bool {
	contentType := t.httpRequest.Header.Get("Content-Type")
	return strings.Contains(contentType, substr)
}

func (t reqStruct) HasAccept(substr string) bool {
	accept := t.httpRequest.Header.Get("Accept")
	return strings.Contains(accept, substr)
}

// Check if the request is form
func (t reqStruct) IsForm() bool {
	return t.HasContentType("application/x-www-form-urlencoded")
}

// Check if the request is json
func (t reqStruct) IsJson() bool {
	return t.HasContentType("json")
}

// Check if the request is multipart
func (t reqStruct) IsMultipart() bool {
	return t.HasContentType("multipart")
}

func (t reqStruct) WantsJson() bool {
	return t.HasAccept("json")
}
