package response

import (
	"bytes"
	html "html/template"
	text "text/template"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/path"
)

type Response interface {
	View(filepaths []string, data interface{})
	Json(data interface{}, status int)
}

// render the templates as string
func Render(filepaths []string, data interface{}) (string, error) {
	for idx, filepath := range filepaths {
		filepaths[idx] = path.Load().ViewPath(filepath)
	}

	t, err := text.ParseFiles(filepaths...)
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

func HTML(filepaths []string, data interface{}) (html.HTML, error) {
	rendered, err := Render(filepaths, data)
	if err != nil {
		return html.HTML(""), err
	}
	return html.HTML(rendered), nil
}