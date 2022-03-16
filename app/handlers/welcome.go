package handlers

import (
	"fmt"
	"github.com/daison12006013/gorvel/pkg/helpers"
	"github.com/daison12006013/gorvel/pkg/storage"
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
	storage := storage.Store

	files, err := req.GetFiles()
	if err != nil {
		return res.Json(helpers.MP{
			"error": err.Error(),
		}, http.StatusOK)
	} // prepare the data

	images := files["files"]

	fmt.Println(len(images))
	for _, image := range images {
		err := storage.Put(image.Filename, image)
		if err != nil {
			return nil
		}
	}

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
		"file":      len(images),
	}
	return res.Json(data, http.StatusOK)

}
