package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/urls"
	"github.com/lucidfy/lucid/pkg/rules"
	"github.com/lucidfy/lucid/pkg/rules/must"
)

type NetHttpRequest struct {
	ResponseWriter     http.ResponseWriter
	HttpRequest        *http.Request
	Translation        *lang.Translations
	URL                *urls.NetHttpURL
	MaxMultipartMemory int64

	ParsedParams map[string]interface{}
}

func NetHttp(w http.ResponseWriter, r *http.Request, t *lang.Translations, u *urls.NetHttpURL) *NetHttpRequest {
	n := NetHttpRequest{
		ResponseWriter:     w,
		HttpRequest:        r,
		Translation:        t,
		URL:                u,
		MaxMultipartMemory: 32 << 20, // 32MB
	}
	return &n
}

type contextKey int

const (
	VarsKey contextKey = iota
)

func (t NetHttpRequest) Vars() map[string]string {
	if rv := t.HttpRequest.Context().Value(VarsKey); rv != nil {
		return rv.(map[string]string)
	}
	return mux.Vars(t.HttpRequest)
}

// All returns available http queries
func (t *NetHttpRequest) All() interface{} {
	// put a singleton, we don't need to parse
	// the request over and over again.
	if len(t.ParsedParams) > 0 {
		return t.ParsedParams
	}

	params := map[string]interface{}{}

	// initialize inputs from route params
	for idx, val := range t.Vars() {
		params[idx] = val
	}

	if t.IsForm() { // via form inputs
		for idx, val := range t.HttpRequest.Form {
			if len(val) > 0 {
				params[idx] = val[0]
			}
		}
	} else if t.IsJson() { // via raw body
		body, err := ioutil.ReadAll(t.HttpRequest.Body)
		if err == nil {
			jsonB := map[string]interface{}{}
			json.Unmarshal(body, &jsonB)
			for idx, val := range jsonB {
				if len(val.(string)) > 0 {
					params[idx] = val
				}
			}
		}
	}

	// get the url queries
	for idx, val := range t.HttpRequest.URL.Query() {
		if len(val) > 0 {
			params[idx] = val[0]
		}
	}

	t.ParsedParams = params

	return params
}

// Get returns the specific value from the provided key
func (t *NetHttpRequest) Get(k string) interface{} {
	// check the queries if exists
	val, ok := t.All().(map[string]interface{})[k]
	if ok {
		return val
	}
	return nil
}

// GetFirst returns the specifc value provided with default value
func (t *NetHttpRequest) GetFirst(k string, dfault interface{}) interface{} {
	val := t.Get(k)
	if val == nil {
		return dfault
	}
	return val
}

// Input ist meant as proxy for GetFirst(...)
func (t *NetHttpRequest) Input(k string, dfault interface{}) interface{} {
	return t.GetFirst(k, dfault)
}

// HasContentType checks if the string exists in the header
func (t *NetHttpRequest) HasContentType(substr string) bool {
	contentType := t.HttpRequest.Header.Get("Content-Type")
	return strings.Contains(contentType, substr)
}

func (t *NetHttpRequest) HasAccept(substr string) bool {
	accept := t.HttpRequest.Header.Get("Accept")
	return strings.Contains(accept, substr)
}

// IsForm checks if the request is an http form
func (t *NetHttpRequest) IsForm() bool {
	return t.HasContentType("application/x-www-form-urlencoded")
}

// IsJson checks if the content type contains json
func (t *NetHttpRequest) IsJson() bool {
	return t.HasContentType("json")
}

// IsMultipart checks if the content type contains multipart
func (t *NetHttpRequest) IsMultipart() bool {
	return t.HasContentType("multipart")
}

func (t *NetHttpRequest) WantsJson() bool {
	return t.HasAccept("json")
}

// Validator
func (t *NetHttpRequest) Validator(setOfRules *must.SetOfRules) *errors.AppError {
	inputValues := t.All().(map[string]interface{})

	validationErrors := rules.
		New(t.Translation, inputValues).
		GetErrors(setOfRules)

	if len(validationErrors) > 0 {
		return &errors.AppError{
			Error:           fmt.Errorf("pkg.facade.request.NetHttprequest@Validator: Request validation error"),
			Message:         "Request validation error",
			Code:            http.StatusUnprocessableEntity,
			ValidationError: validationErrors,
		}
	}

	return nil
}

// GetIp returns the client IP address
// it resolves first from "x-forwarded-for"
// or else it goes check if "x-real-ip" exists
// or else we pull based on the remoteaddr under net/http
func (t *NetHttpRequest) GetIp() string {
	ip := t.HttpRequest.Header.Get("X-Forwarded-For")
	if len(ip) == 0 {
		ip = t.HttpRequest.Header.Get("X-Real-Ip")
	}
	if len(ip) == 0 {
		ip = t.HttpRequest.RemoteAddr
	}
	return ip
}

// GetUserAgent returns the user agent
func (t *NetHttpRequest) GetUserAgent() string {
	return t.HttpRequest.Header.Get("User-Agent")
}

// GetFileByName returns the first file for the provided form key.
func (t *NetHttpRequest) GetFileByName(name string) (*multipart.FileHeader, *errors.AppError) {
	if t.HttpRequest.MultipartForm == nil {
		if err := t.HttpRequest.ParseMultipartForm(t.MaxMultipartMemory); err != nil {
			return nil, errors.InternalServerError("t.HttpRequest.ParseMultipartForm() error", err)
		}
	}

	f, fh, err := t.HttpRequest.FormFile(name)
	if err != nil {
		return nil, errors.InternalServerError("t.HttpRequest.FormFile() error", err)
	}

	err = f.Close()

	return fh, errors.InternalServerError("f.Close() error", err)
}

// GetFiles is the parsed multipart form files
func (t *NetHttpRequest) GetFiles() (map[string][]*multipart.FileHeader, *errors.AppError) {
	err := t.HttpRequest.ParseMultipartForm(t.MaxMultipartMemory)
	return t.HttpRequest.MultipartForm.File, errors.InternalServerError("t.HttpRequest.ParseMultipartForm() error", err)
}
