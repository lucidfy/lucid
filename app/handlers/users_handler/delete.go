package users_handler

import (
	"net/http"

	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid"
)

func delete(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	//> prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	//> validate "id" if exists
	id := req.Input("id", nil).(string)
	if app_err := users.Exists("id", &id); app_err != nil {
		return app_err
	}

	//> now get the data
	data, app_err := users.Find(&id, nil)
	if app_err != nil {
		return app_err
	}

	//> and delete the data
	data.Delete()

	//> response: for api based
	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
		}, status)
	}

	//> response: for form based, just redirect
	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
