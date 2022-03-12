package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/path"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func Docs(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	r := engine.HttpRequest
	res := engine.Response

	//> detect the url path, just that we replace any suffix that has /docs
	// then we fetch the remaining file name
	f := strings.Replace(r.URL.Path, "/docs", "", -1)
	if len(f) == 0 {
		f = "index"
	}

	title := strings.Trim(f, "/")

	//> let's make sure the file contains .md format, or else append it
	if !strings.Contains(f, ".md") {
		f = f + ".md"
	}

	//> let's read the full path of the file as markdown file.
	d := http.Dir(path.PathTo("/resources/docs/" + f))
	md, err := os.ReadFile(string(d))
	if err != nil {
		return &errors.AppError{
			Error: err,
		}
	}

	flags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Title: "A custom title",
		Flags: flags,
	}
	renderer := html.NewRenderer(opts)
	asHtml := markdown.ToHTML(md, nil, renderer)

	return res.View(
		[]string{"base", "docs"},
		map[string]interface{}{
			"title": title,
			"md":    string(asHtml),
		},
	)
}
