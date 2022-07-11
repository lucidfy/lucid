package response

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	text "text/template"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/path"
)

type NetHttpResponse struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func NetHttp(w http.ResponseWriter, r *http.Request) *NetHttpResponse {
	return &NetHttpResponse{
		ResponseWriter: w,
		HttpRequest:    r,
	}
}

func (m NetHttpResponse) Default() http.ResponseWriter {
	return m.ResponseWriter
}

func (m *NetHttpResponse) Text(str string) *errors.AppError {
	io.WriteString(m.ResponseWriter, str)
	return nil
}

func (m *NetHttpResponse) ViewWithStatus(filepaths []string, data interface{}, status int) *errors.AppError {
	m.ResponseWriter.Header().Set("Content-Type", "text/html")
	m.ResponseWriter.WriteHeader(status)

	for idx, filepath := range filepaths {
		if !strings.Contains(filepath, DEFAULT_VIEW_EXT) {
			filepath = filepath + DEFAULT_VIEW_EXT
		}

		filepaths[idx] = path.Load().ViewPath(filepath)
	}

	data = m.constructDataFromHeader(
		data,
		m.ResponseWriter.Header().Get("X-CSRF-Token"),
		"csrf_token",
	)

	t, err := text.ParseFiles(filepaths...)
	if err != nil {
		return errors.InternalServerError("Error text.ParseFiles()", err)
	}

	err = t.Execute(m.ResponseWriter, data)
	if err != nil {
		return errors.InternalServerError("Error template.Execute()", err)
	}

	return nil
}

func (m *NetHttpResponse) View(filepaths []string, data interface{}) *errors.AppError {
	return m.ViewWithStatus(filepaths, data, http.StatusOK)
}

func (m *NetHttpResponse) Json(data interface{}, status int) *errors.AppError {
	m.ResponseWriter.Header().Set("Content-Type", "application/json")
	m.ResponseWriter.WriteHeader(status)

	err := json.NewEncoder(m.ResponseWriter).Encode(data)
	if err != nil {
		return errors.InternalServerError("Error encoding json data", err)
	}

	return nil
}

func (m *NetHttpResponse) constructDataFromHeader(data interface{}, val string, key string) interface{} {
	if len(val) > 0 {
		if m, ok := (data).(map[string]interface{}); ok {
			m[key] = val
			data = m
		}
	}
	return data
}
