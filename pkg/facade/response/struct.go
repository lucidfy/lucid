package response

import (
	"bytes"
	html "html/template"
	"strings"
	text "text/template"

	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/logger"
	"github.com/daison12006013/lucid/pkg/facade/path"
)

const DEFAULT_VIEW_EXT = ".go.html"

type ResponseContract interface {
	ViewWithStatus(filepaths []string, data interface{}, status *int) *errors.AppError
	View(filepaths []string, data interface{}) *errors.AppError
	Json(data interface{}, status int) *errors.AppError
}

// render the templates as string
func Render(filepaths []string, data interface{}) (string, error) {
	for idx, filepath := range filepaths {
		if !strings.Contains(filepath, DEFAULT_VIEW_EXT) {
			filepath = filepath + DEFAULT_VIEW_EXT
		}

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
