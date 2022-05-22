package single_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/routes"
)

var WelcomeRoute = routes.Routing{
	Path:    "/",
	Name:    "welcome",
	Method:  routes.Method{"GET"},
	Handler: welcome,
}

func welcome(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// prepare the data
	data := map[string]interface{}{
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
