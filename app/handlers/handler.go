package handlers

import (
	"os"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
)

// HttpErrorHandler,  we handle any returned errors.AppError under this
// we serve html or a json format
func HttpErrorHandler(T engines.EngineContract, app_err *errors.AppError) {
	engine := T.(engines.NetHttpEngine)
	req := engine.Request
	res := engine.Response

	// assign a default message and code here
	code := 500
	message := "Something went wrong!"
	if app_err.Code != nil && app_err.Code.(int) != 0 {
		code = app_err.Code.(int)
	}
	if app_err.Message != nil && len(app_err.Message.(string)) > 0 {
		message = app_err.Message.(string)
	}

	// initialize the data
	data := map[string]interface{}{
		"message": message,
		"code":    code,
	}

	// don't provide the real error if the debug is not true!
	if os.Getenv("APP_DEBUG") == "true" {
		data["error"] = app_err.Error
	}

	// write a json format
	if req.WantsJson() {
		res.Json(data, code)
		return
	}

	// write html format
	res.ViewWithStatus([]string{"pkg/error/default"}, data, code)
}
