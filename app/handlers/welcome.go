package handlers

import (
	"github.com/daison12006013/gorvel/pkg/helpers"
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

func Welcome(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// prepare the data
	data := map[string]interface{}{
		"title":     "Gorvel Rocks! ",
		"IpAddress": req.GetIp(),
		"userAgent": req.GetUserAgent(),
	}

	// this is api request
	if req.IsJson() && req.WantsJson() {
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

func WelcomeForApi(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	res := engine.Response

	file, err := req.GetFileByName("file")
	if err != nil {
		return res.Json(helpers.MP{
			"error": err.Error(),
		}, http.StatusOK)
	}
	// prepare the data
	data := map[string]interface{}{
		"title":     "Gorvel Rocks! ",
		"IpAddress": req.GetIp(),
		"userAgent": req.GetUserAgent(),
		"file":      file.Filename,
	}
	return res.Json(data, http.StatusOK)

}
