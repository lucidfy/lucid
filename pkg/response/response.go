package response

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/path"
)

func View(w http.ResponseWriter, filepaths []string, data interface{}) {
	w.WriteHeader(http.StatusOK)

	for idx, filepath := range filepaths {
		filepaths[idx] = path.Load().ViewPath(filepath)
	}

	// validate that the csrf token exists
	// append the csrf token inside
	csrfToken := w.Header().Get("X-CSRF-Token")
	if len(csrfToken) > 0 {
		if m, ok := (data).(map[string]interface{}); ok {
			m["csrf_token"] = csrfToken
			data = m
		}
	}

	t, err := template.ParseFiles(filepaths...)

	if err != nil {
		logger.Fatal(err)
	}

	t.Execute(w, data)
}

func Json(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
