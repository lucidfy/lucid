package response

import (
	"net/http"
	"text/template"

	"github.com/daison12006013/gorvel/internal/filemanager"
)

func View(w http.ResponseWriter, filepath string, data interface{}) {
	filepath = filemanager.PathTo("/resources/views/") + filepath
	t, _ := template.ParseFiles(filepath)
	t.Execute(w, data)
}
