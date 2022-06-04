package sample_handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/iancoleman/strcase"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/functions/php"
	"github.com/lucidfy/lucid/pkg/lucid"
)

var DocsRoute = routes.Routing{
	Path:    "/docs",
	Prefix:  true,
	Name:    "docs",
	Method:  routes.Method{"GET"},
	Handler: docs,
}

func docs(ctx context.Context) *errors.AppError {
	engine := lucid.Context(ctx).Engine()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	//> detect the url path, just that we replace any suffix that has /docs
	// then we fetch the remaining file name
	f := strings.Replace(url.CurrentURL(), "/docs", "", -1)
	title := strings.Trim(f, "/")
	if len(title) == 0 {
		title = "Lucid"
		f = "index"
	}

	//> let's make sure the file contains .md format, or else append it
	if !strings.Contains(f, ".md") {
		f = f + ".md"
	}

	//> let's read the full path of the file as markdown file.
	d := http.Dir(path.PathTo("/resources/docs/" + f))

	if !php.FileExists(string(d)) {
		return res.Json(nil, http.StatusNotFound)
	}

	ext := parser.CommonExtensions | parser.Attributes | parser.OrderedListStart | parser.SuperSubscript | parser.Mmark
	var asHtml []byte

	if req.Input("parse", 0) == "1" {
		asHtml = markdown.ToHTML(
			*php.FileGetContents(string(d)),
			parser.NewWithExtensions(ext),
			nil,
		)
	} else {
		asHtml = *php.FileGetContents(string(d))
	}

	return res.Json(map[string]interface{}{
		"title":   strings.Title(strcase.ToDelimited(title, ' ')),
		"content": string(asHtml),
	}, http.StatusOK)
}
