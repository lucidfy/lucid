package request

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"

	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/session"
	"github.com/daison12006013/gorvel/pkg/facade/urls"
	"github.com/daison12006013/gorvel/pkg/rules"
	"github.com/daison12006013/gorvel/pkg/rules/must"
	"github.com/gorilla/mux"
)

type MuxRequest struct {
	ResponseWriter     http.ResponseWriter
	HttpRequest        *http.Request
	Url                *urls.MuxUrl
	Session            *session.MuxSession
	MaxMultipartMemory int64
}

func Mux(w http.ResponseWriter, r *http.Request, u *urls.MuxUrl, ses *session.MuxSession) *MuxRequest {
	t := MuxRequest{
		ResponseWriter:     w,
		HttpRequest:        r,
		Url:                u,
		Session:            ses,
		MaxMultipartMemory: 32 << 20, // 32MB
	}
	return &t
}

// All returns available http queries
func (t *MuxRequest) All() interface{} {
	params := map[string]interface{}{}

	// via form inputs
	for idx, val := range t.HttpRequest.Form {
		if len(val) > 0 {
			params[idx] = val[0]
		}
	}

	// via query params
	for idx, val := range t.HttpRequest.URL.Query() {
		if len(val) > 0 {
			params[idx] = val[0]
		}
	}

	// via route params
	for idx, val := range mux.Vars(t.HttpRequest) {
		params[idx] = val
	}

	return params
}

// Get returns the specific value from the provided key
func (t *MuxRequest) Get(k string) interface{} {
	// check the queries if exists
	val, ok := t.All().(map[string]interface{})[k]
	if ok {
		return val
	}
	return nil
}

// GetFirst returns the specifc value provided with default value
func (t *MuxRequest) GetFirst(k string, dfault interface{}) interface{} {
	val := t.Get(k)
	if val == nil {
		return dfault
	}
	return val
}

// Input ist meant as proxy for GetFirst(...)
func (t *MuxRequest) Input(k string, dfault interface{}) interface{} {
	return t.GetFirst(k, dfault)
}

// HasContentType checks if the string exists in the header
func (t *MuxRequest) HasContentType(substr string) bool {
	contentType := t.HttpRequest.Header.Get("Content-Type")
	return strings.Contains(contentType, substr)
}

func (t *MuxRequest) HasAccept(substr string) bool {
	accept := t.HttpRequest.Header.Get("Accept")
	return strings.Contains(accept, substr)
}

// IsForm checks if the request is an http form
func (t *MuxRequest) IsForm() bool {
	return t.HasContentType("application/x-www-form-urlencoded")
}

// IsJson checks if the content type contains json
func (t *MuxRequest) IsJson() bool {
	return t.HasContentType("json")
}

// IsMultipart checks if the content type contains multipart
func (t *MuxRequest) IsMultipart() bool {
	return t.HasContentType("multipart")
}

func (t *MuxRequest) WantsJson() bool {
	return t.HasAccept("json")
}

// --- Validator

func (t *MuxRequest) Validator(setOfRules *must.SetOfRules) *errors.AppError {
	var errsChan = make(chan map[string]string)

	var wg sync.WaitGroup

	for inputField, inputRules := range *setOfRules {
		for _, inputRule := range inputRules {
			inputValue := t.Get(inputField)
			wg.Add(1)
			go rules.Validate(
				inputField,
				fmt.Sprint(inputValue),
				inputRule,
				errsChan,
				&wg,
			)
		}
	}

	go func() {
		wg.Wait()
		close(errsChan)
	}()

	validationErrors := map[string]interface{}{}
	for val := range errsChan {
		for k, v := range val {
			validationErrors[k] = v
		}
	}

	if len(validationErrors) > 0 {
		return &errors.AppError{
			Error:           fmt.Errorf("pkg.facade.request.muxrequest@Validator: Request validation error"),
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
func (t *MuxRequest) GetIp() string {
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
func (t *MuxRequest) GetUserAgent() string {
	return t.HttpRequest.Header.Get("User-Agent")
}

// GetFileByName returns the first file for the provided form key.
func (t *MuxRequest) GetFileByName(name string) (*multipart.FileHeader, error) {
	if t.HttpRequest.MultipartForm == nil {
		if err := t.HttpRequest.ParseMultipartForm(t.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := t.HttpRequest.FormFile(name)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return fh, err
}
