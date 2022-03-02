package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

func Create(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	// req := engine.Request
	res := engine.Response

	data := map[string]interface{}{"title": "Create Form"}

	return res.View(
		[]string{"base", "users/create"},
		data,
	)
}

func Store(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	// req := engine.Request
	res := engine.Response

	// prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	// * TODO

	// prepare the data
	return res.Json(map[string]interface{}{
		"message": message,
	}, status)
}
