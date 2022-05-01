package authhandler

import (
	"net/http"

	"github.com/daison12006013/lucid/app/models/users"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/hash"
	"github.com/daison12006013/lucid/pkg/facade/session"
)

func User(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	res := engine.Response

	userID, err := ses.Get("authenticated")
	if userID != nil || err != nil {
		return res.Json(map[string]interface{}{}, http.StatusOK)
	}

	data, appErr := users.Find(userID, nil)
	if appErr != nil {
		return appErr
	}

	return res.Json(map[string]interface{}{
		"user": data.Model,
	}, http.StatusOK)
}

func LoginAttempt(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	email := req.Input("email", nil).(string)
	password := req.Input("password", nil).(string)

	if appErr := users.Exists("email", &email); appErr != nil {
		appErr.Message = "Email or Password is incorrect!"
		return appErr
	}

	data, appErr := users.Find(&email, "email")
	if appErr != nil {
		appErr.Message = "Email or Password is incorrect!"
		return appErr
	}

	record := data.Model

	if hash.Check(password, record.Password) {
		ses.Set("authenticated", record.ID)
	}

	message := "Successfully Logged In!"
	status := http.StatusOK

	if req.WantsJson() && req.IsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    record,
		}, status)
	}

	ses.SetFlash("success", message)
	url.RedirectPrevious()
	return nil
}
