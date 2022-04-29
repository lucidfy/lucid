package response

import (
	"encoding/json"
	"net/http"
	"strings"
	text "text/template"

	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/path"
)

type MuxResponse struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) *MuxResponse {
	m := MuxResponse{
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return &m
}

func (m *MuxResponse) ViewWithStatus(filepaths []string, data interface{}, status *int) *errors.AppError {
	m.ResponseWriter.WriteHeader(*status)

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
		return &errors.AppError{
			Error:   err,
			Message: "Error parsing files",
			Code:    http.StatusInternalServerError,
		}
	}

	err = t.Execute(m.ResponseWriter, data)
	if err != nil {
		return &errors.AppError{
			Error:   err,
			Message: "Error executing template",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (m *MuxResponse) View(filepaths []string, data interface{}) *errors.AppError {
	httpOk := 200
	return m.ViewWithStatus(filepaths, data, &httpOk)
}

func (m *MuxResponse) Json(data interface{}, status int) *errors.AppError {
	m.ResponseWriter.Header().Set("Content-Type", "application/json")
	m.ResponseWriter.WriteHeader(status)
	err := json.NewEncoder(m.ResponseWriter).Encode(data)
	if err != nil {
		return &errors.AppError{
			Error:   err,
			Message: "Error encoding json data",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func (m *MuxResponse) constructDataFromHeader(data interface{}, val string, key string) interface{} {
	if len(val) > 0 {
		if m, ok := (data).(map[string]interface{}); ok {
			m[key] = val
			data = m
		}
	}
	return data
}
