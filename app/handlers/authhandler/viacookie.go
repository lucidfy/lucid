package authhandler

import (
	"net/http"

	"github.com/daison12006013/lucid/app/models/users"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/hash"
)

func ViaCookie(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	req := engine.Request
	res := engine.Response
	ses := engine.Session
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
		ses.Set("user", record.ID)
	}

	message := "Successfully Created!"
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
