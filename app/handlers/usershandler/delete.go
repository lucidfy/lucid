package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
)

func Delete(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response
	ses := engine.Session
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
	data, err := users.Find(&id)
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
