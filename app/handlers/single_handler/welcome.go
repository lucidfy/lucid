package single_handler

import (
	"context"
	"net/http"

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

func welcome(ctx context.Context) *errors.AppError {
	engine := lucid.Context(ctx).Engine()
	req := engine.GetRequest()
	res := engine.GetResponse()

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
