package response

import (
	"encoding/json"
	"net/http"
	text "text/template"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/path"
)

type MuxResponse struct {
	ResponseWriter http.ResponseWriter
	HttpRequest    *http.Request
}

func Mux(w http.ResponseWriter, r *http.Request) MuxResponse {
	t := MuxResponse{
		ResponseWriter: w,
		HttpRequest:    r,
	}
	return t
}

func (m MuxResponse) View(filepaths []string, data interface{}) {
	m.ResponseWriter.WriteHeader(http.StatusOK)

	for idx, filepath := range filepaths {
		filepaths[idx] = path.Load().ViewPath(filepath)
	}

	data = m.constructDataFromHeader(
		data,
		m.ResponseWriter.Header().Get("X-CSRF-Token"),
		"csrf_token",
	)

	t, err := text.ParseFiles(filepaths...)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	t.Execute(m.ResponseWriter, data)
}

func (m MuxResponse) constructDataFromHeader(data interface{}, val string, key string) interface{} {
	if len(val) > 0 {
		if m, ok := (data).(map[string]interface{}); ok {
			m[key] = val
			data = m
		}
	}
	return data
}

func (m MuxResponse) Json(data interface{}, status int) {
	m.ResponseWriter.Header().Set("Content-Type", "application/json")
	m.ResponseWriter.WriteHeader(status)
	json.NewEncoder(m.ResponseWriter).Encode(data)
}
