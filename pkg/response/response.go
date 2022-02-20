package response

import (
	"bytes"
	"encoding/json"
	h "html/template"
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

	data = constructDataFromHeader(data, w.Header().Get("X-CSRF-Token"), "csrf_token")

	t, err := template.ParseFiles(filepaths...)
	if err != nil {
		logger.Fatal(err)
		panic(err)
	}

	t.Execute(w, data)
}

func constructDataFromHeader(data interface{}, val string, key string) interface{} {
	if len(val) > 0 {
		if m, ok := (data).(map[string]interface{}); ok {
			m[key] = val
			data = m
		}
	}
	return data
}

func Json(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// render the templates as string
func Render(filepaths []string, data interface{}) (string, error) {
	for idx, filepath := range filepaths {
		filepaths[idx] = path.Load().ViewPath(filepath)
	}

	t, err := template.ParseFiles(filepaths...)
	if err != nil {
		logger.Fatal(err)
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		logger.Fatal(err)
		return "", err
	}

	return tpl.String(), nil
}

func HTML(filepaths []string, data interface{}) (h.HTML, error) {
	rendered, err := Render(filepaths, data)
	if err != nil {
		return h.HTML(""), err
	}
	return h.HTML(rendered), nil
}
