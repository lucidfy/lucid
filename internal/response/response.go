package response

import (
	"net/http"
	"text/template"

	"github.com/daison12006013/gorvel/internal/facade/path"
)

func View(w http.ResponseWriter, filepath string, data interface{}) {
	filepath = path.Load().ViewPath(filepath)
	t, _ := template.ParseFiles(filepath)
	t.Execute(w, data)
}
