package single_handler

import (
	"net/http"
	"os"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/lucid"
)

var WelcomeRoute = routes.Routing{
	Path:    "/",
	Name:    "welcome",
	Method:  routes.Method{"GET"},
	Handler: welcome,
}

func welcome(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	req := engine.GetRequest()
	res := engine.GetResponse()
	lang := engine.GetTranslation().SetLanguage(
		req.Input("language", os.Getenv("APP_LANGUAGE")).(string),
	)

	// prepare the data
	data := map[string]interface{}{
		"lang":  lang,
		"title": "Lucid Rocks!",
	}

	// this is api request
	if req.WantsJson() {
		return res.Json(data, http.StatusOK)
	}

	// render the template
	return res.View(
		// this example below, we're telling the compiler
		// to parse the base.go.html first, and then parse the welcome.go.html
		// therefore the defined "body" should render accordingly
		[]string{"base", "welcome"},
		data,
	)
}
