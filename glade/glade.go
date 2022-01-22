package glade

import (
	"net/http"
	"text/template"
)

func View(w http.ResponseWriter, file string, data interface{}) {
	t, _ := template.ParseFiles(file)
	t.Execute(w, data)
}
