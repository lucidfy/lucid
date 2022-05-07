package handlers

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
)

func Welcome(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// prepare the data
	data := map[string]interface{}{
		"title":     "Lucid Rocks! ",
		"IpAddress": req.GetIp(),
		"userAgent": req.GetUserAgent(),
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
