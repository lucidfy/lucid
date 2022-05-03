package handlers

import (
	"net/http"
	"strings"

	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/path"
	"github.com/daison12006013/lucid/pkg/functions/php"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func Docs(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	//> detect the url path, just that we replace any suffix that has /docs
	// then we fetch the remaining file name
	f := strings.Replace(r.URL.Path, "/docs", "", -1)
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
	asHtml := markdown.ToHTML(
		*php.FileGetContents(string(d)),
		parser.NewWithExtensions(ext),
		nil,
	)

	data := map[string]interface{}{
		"title":   title,
		"content": string(asHtml),
	}

	if req.WantsJson() && req.IsJson() {
		return res.Json(data, http.StatusOK)
	}

	return res.View(
		[]string{"base", "docs"},
		data,
	)
}
