package view

import (
	"net/http"
	"text/template"
)

func Render(w http.ResponseWriter, file string, data interface{}) {
	t, _ := template.ParseFiles(file)
	t.Execute(w, data)
}
