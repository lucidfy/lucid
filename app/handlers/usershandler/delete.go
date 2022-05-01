package usershandler

import (
	"net/http"

	"github.com/daison12006013/lucid/app/models/users"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/session"
)

func Delete(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	//> prepare message and status
	message := "Successfully Deleted!"
	status := http.StatusOK

	//> validate "id" if exists
	id := req.Input("id", nil).(string)
	if err := users.Exists("id", &id); err != nil {
		return err
	}

	//> now get the data
	data, err := users.Find(&id, nil)
	if err != nil {
		return err
	}

	//> and delete the data
	data.Delete()

	//> response: for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
		}, status)
	}

	//> response: for form based, just redirect
	ses.SetFlash("success", message)
	url.RedirectPrevious()
	return nil
}
