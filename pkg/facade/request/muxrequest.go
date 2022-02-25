package request

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/facade/urls"
	"github.com/gorilla/mux"
)

type MuxRequest struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) MuxRequest {
	t := MuxRequest{
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return t
}

// Url  ----------------------------------------------------

func (t MuxRequest) CurrentUrl() string {
	h := t.HttpRequest.URL.Host
	if len(h) > 0 {
		return h
	}
	// if ever we can't resolve the Host from net/http, we can still base
	// from our env config
	return urls.BaseUrl(nil)
}

func (t MuxRequest) FullUrl() string {
	return t.CurrentUrl() + t.HttpRequest.URL.RequestURI()
}

func (t MuxRequest) PreviousUrl() string {
	return t.HttpRequest.Referer()
}

func (t MuxRequest) RedirectPrevious() {
	http.Redirect(t.ResponseWriter, t.HttpRequest, t.PreviousUrl(), http.StatusFound)
}

func (t MuxRequest) SetFlash(name string, value string) {
	name = "flash-" + name
	s := session.Mux(t.ResponseWriter, t.HttpRequest)
	s.Set(name, value)
}

func (t MuxRequest) GetFlash(name string) *string {
	name = "flash-" + name

	s := session.Mux(t.ResponseWriter, t.HttpRequest)
	if s == nil {
		return nil
	}

	value, err := s.Get(name)
	if (err != nil && err == http.ErrNoCookie) || value == nil {
		return nil
	}

	// delete the cookie by expiring it immediately!
	deleteCookie := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0), Path: "/"}
	http.SetCookie(t.ResponseWriter, deleteCookie)
	return value
}

// Request  -------------------------------------------------

// This returns all avaiable queries from
func (t MuxRequest) All() url.Values {
	params := t.HttpRequest.URL.Query()
	return params
}

// This returns the specific value from the provided key
func (t MuxRequest) Get(k string) []string {
	// check the queries if exists
	val, ok := t.All()[k]
	if ok {
		return val
	}

	// or try to extract from the route
	vars := mux.Vars(t.HttpRequest)
	if val, ok := vars[k]; ok {
		return []string{val}
	}

	return []string{}
}

func (t MuxRequest) GetFirst(k string, dfault *string) *string {
	val := t.Get(k)
	if len(val) > 0 {
		return &val[0]
	}
	return dfault
}

// Proxy method to GetFirst(...)
func (t MuxRequest) Input(k string, dfault string) string {
	d := t.GetFirst(k, &dfault)
	if d == nil {
		return ""
	}
	return *d
}

// Check if the string exists in the content type
func (t MuxRequest) HasContentType(substr string) bool {
	contentType := t.HttpRequest.Header.Get("Content-Type")
	return strings.Contains(contentType, substr)
}

func (t MuxRequest) HasAccept(substr string) bool {
	accept := t.HttpRequest.Header.Get("Accept")
	return strings.Contains(accept, substr)
}

// Check if the request is form
func (t MuxRequest) IsForm() bool {
	return t.HasContentType("application/x-www-form-urlencoded")
}

// Check if the request is json
func (t MuxRequest) IsJson() bool {
	return t.HasContentType("json")
}

// Check if the request is multipart
func (t MuxRequest) IsMultipart() bool {
	return t.HasContentType("multipart")
}

func (t MuxRequest) WantsJson() bool {
	return t.HasAccept("json")
}
