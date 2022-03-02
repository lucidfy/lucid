package usershandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/gorilla/csrf"
)

func Show(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	request := engine.Request
	response := engine.Response

	id := request.Input("id", nil).(string)
	if err := users.Exists(&id); err != nil {
		return err
	}

	record := users.Find(&id).Record
	if request.IsJson() && request.WantsJson() {
		return response.Json(record, http.StatusOK)
	}

	// by default we use "show" then check if the
	// url path contains /edit , therefore use "edit"
	viewFile := "show"
	if strings.Contains(engine.HttpRequest.URL.Path, "/edit") {
		viewFile = "edit"
	}

	return response.View(
		[]string{"base", fmt.Sprintf("users/%s", viewFile)},
		map[string]interface{}{
			"title":          record.Name + "'s Profile",
			"previousUrl":    request.PreviousUrl(),
			"record":         record,
			"success":        request.GetFlash("success"),
			"error":          request.GetFlash("error"),
			csrf.TemplateTag: csrf.TemplateField(engine.HttpRequest),
		},
	)
}

func Update(T engines.EngineInterface) *errors.AppError {
	engine := T.(engines.MuxEngine)
	// w := engine.HttpResponseWriter
	// r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	message := "Successfully Updated!"
	status := http.StatusOK

	id := req.Input("id", nil).(string)
	found := users.Find(&id)
	found.Updates(req.All())

	// for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
	}

	// for form based, just redirect
	req.SetFlash("success", message)
	req.RedirectPrevious()
	return nil
}
