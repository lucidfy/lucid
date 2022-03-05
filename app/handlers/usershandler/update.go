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
	// w := engine.HttpResponseWriter
	r := engine.HttpRequest
	req := engine.Request
	res := engine.Response

	id := req.Input("id", nil).(string)
	if appErr := users.Exists("id", &id); appErr != nil {
		return appErr
	}

	data, appErr := users.Find(&id)
	if appErr != nil {
		return appErr
	}

	record := data.Model
	if req.IsJson() && req.WantsJson() {
		return res.Json(record, http.StatusOK)
	}

	// by default we use "show" then check if the
	// url path contains /edit , therefore use "edit"
	isShow := true
	viewFile := "show"
	if strings.Contains(r.URL.Path, "/edit") {
		isShow = false
	}

	return res.View(
		[]string{"base", fmt.Sprintf("users/%s", viewFile)},
		map[string]interface{}{
			"title":          record.Name + "'s Profile",
			"previousUrl":    req.PreviousUrl(),
			"record":         record,
			"isShow":         isShow,
			csrf.TemplateTag: csrf.TemplateField(r),

			"success": req.GetFlash("success"),
			"error":   req.GetFlash("error"),
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
	data, appErr := users.Find(&id)
	if appErr != nil {
		return appErr
	}
	data.Updates(req.All())

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
