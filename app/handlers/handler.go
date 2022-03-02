// please avoid deleting this handler
// this handler is meant to show a 404 and 405 page
// if ever a user navigated into a wrong url page
// nor submitting a wrong method into registered route
package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

func ModelBinding(T engines.EngineInterface) {
	// engine := T.(engines.MuxEngine)
}

// on this method, once the engine detects a page not found
// (404 Code) requests, it should divert back to this handler
func PageNotFound(T engines.EngineInterface) {
	err := &errors.AppError{
		Code:    http.StatusNotFound,
		Message: "Page not found",
		Error:   fmt.Errorf("404 page not found"),
	}
	HttpErrorHandler(T, err)
}

// on this method, once the engine detects a method not allowed
// (405 Code) requests, it should divert back to this handler
func MethodNotAllowed(T engines.EngineInterface) {
	err := &errors.AppError{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
		Error:   fmt.Errorf("405 method not allowed"),
	}
	HttpErrorHandler(T, err)
}

// on this method, we handle the any returned AppError across all handlers
// once we received any errors.AppError, what we actually do is to print
// a pretty neat html (or if the requestor wanted a json, we respond accordingly)
func HttpErrorHandler(T engines.EngineInterface, appErr *errors.AppError) {
	engine := T.(engines.MuxEngine)

	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	// assign a default values here
	code := 500
	message := "Something went wrong!"
	if appErr.Code.(int) != 0 {
		code = appErr.Code.(int)
	}
	if len(appErr.Message.(string)) > 0 {
		message = appErr.Message.(string)
	}

	// initialize the data
	data := map[string]interface{}{
		"message": message,
		"code":    code,
	}

	// don't provide the real error if the debug is not true!
	if os.Getenv("APP_DEBUG") == "true" {
		data["error"] = appErr.Error
	}

	// write a json format
	if req.IsJson() && req.WantsJson() {
		res.Json(data, code)
		return
	}

	// write html format
	res.ViewWithStatus([]string{"pkg/error/default"}, data, &code)
}
