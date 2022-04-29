package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/path"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func Docs(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	r := engine.HttpRequest
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
	md, err := os.ReadFile(string(d))
	if err != nil {
		return &errors.AppError{
			Error: err,
		}
	}

	ext := parser.CommonExtensions | parser.Attributes | parser.OrderedListStart | parser.SuperSubscript | parser.Mmark
	asHtml := markdown.ToHTML(
		md,
		parser.NewWithExtensions(ext),
		nil,
	)

	return res.View(
		[]string{"base", "docs"},
		map[string]interface{}{
			"title": title,
			"md":    string(asHtml),
			"menus": menus(),
		},
	)
}

type MenuAttr struct {
	Name string
	URL  string
}
type MenuChildren struct {
	HasChild bool
	Children []MenuAttr
}
type Menu struct {
	MenuAttr
	MenuChildren
}

func menus() *[]Menu {
	return &[]Menu{
		{
			MenuAttr{
				Name: "Prologue",
				URL:  "",
			},
			MenuChildren{
				HasChild: true,
				Children: []MenuAttr{
					{
						Name: "Contribution Guide",
						URL:  "/docs/Contribution Guide",
					},
				},
			},
		},
		{
			MenuAttr{
				Name: "Getting Started",
				URL:  "",
			},
			MenuChildren{
				HasChild: true,
				Children: []MenuAttr{
					{
						Name: "Installation",
						URL:  "/docs/Installation",
					},
				},
			},
		},
		{
			MenuAttr{
				Name: "Core Documentation",
				URL:  "https://pkg.go.dev/github.com/daison12006013/lucid",
			},
			MenuChildren{},
		},
	}
}
