// Package handlers > please avoid deleting this handler
//> this handler is meant to show a 404 and 405 page
//> if ever a user navigated into a wrong url page
//> nor submitting a wrong method into registered route
package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
)

// PageNotFound > on this method, once the engine detects a page not found
//> (404 Code) requests, it should divert back to this handler
func PageNotFound(T engines.EngineContract) {
	app_err := &errors.AppError{
		Code:    http.StatusNotFound,
		Message: "Page not found",
		Error:   fmt.Errorf("404 page not found"),
	}
	HttpErrorHandler(T, app_err)
}

// MethodNotAllowed > on this method, once the engine detects a method not allowed
//> (405 Code) requests, it should divert back to this handler
func MethodNotAllowed(T engines.EngineContract) {
	err := &errors.AppError{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method not allowed",
		Error:   fmt.Errorf("405 method not allowed"),
	}
	HttpErrorHandler(T, err)
}

// HttpErrorHandler > on this method, we handle the any returned AppError across all handlers
//> once we received any errors.AppError, what we actually do is to print
//> a pretty neat html (or if the requestor wanted a json, we respond accordingly)
func HttpErrorHandler(T engines.EngineContract, app_err *errors.AppError) {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	//> assign a default message and code here
	code := 500
	message := "Something went wrong!"
	if app_err.Code != nil && app_err.Code.(int) != 0 {
		code = app_err.Code.(int)
	}
	if app_err.Message != nil && len(app_err.Message.(string)) > 0 {
		message = app_err.Message.(string)
	}

	//> initialize the data
	data := map[string]interface{}{
		"message": message,
		"code":    code,
	}

	//> don't provide the real error if the debug is not true!
	if os.Getenv("APP_DEBUG") == "true" {
		data["error"] = app_err.Error
	}

	//> write a json format
	if req.WantsJson() {
		res.Json(data, code)
		return
	}

	//> write html format
	res.ViewWithStatus([]string{"pkg/error/default"}, data, &code)
}
