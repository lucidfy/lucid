package auth_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/hash"
)

func user(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.ResponseWriter
	// r := engine.HttpRequest
	ses := engine.Session
	res := engine.Response

	userID, app_err := ses.Get("authenticated")
	if userID != nil || app_err != nil {
		return res.Json(map[string]interface{}{}, http.StatusOK)
	}

	data, app_err := users.Find(userID, nil)
	if app_err != nil {
		return app_err
	}

	return res.Json(map[string]interface{}{
		"user": data.Model,
	}, http.StatusOK)
}

func loginAttempt(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.NetHttpEngine)
	// w := engine.ResponseWriter
	// r := engine.HttpRequest
	ses := engine.Session
	req := engine.Request
	res := engine.Response
	url := engine.URL

	email := req.Input("email", nil).(string)
	password := req.Input("password", nil).(string)

	if app_err := users.Exists("email", &email); app_err != nil {
		app_err.Message = "Email or Password is incorrect!"
		return app_err
	}

	data, app_err := users.Find(&email, "email")
	if app_err != nil {
		app_err.Message = "Email or Password is incorrect!"
		return app_err
	}

	record := data.Model

	if hash.Check(password, record.Password) {
		ses.Put("authenticated", record.ID)
	}

	message := "Successfully Logged In!"
	status := http.StatusOK

	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    record,
		}, status)
	}

	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
