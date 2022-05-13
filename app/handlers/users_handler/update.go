package users_handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/session"
)

func Show(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	id := req.Input("id", nil).(string)
	if appErr := users.Exists("id", &id); appErr != nil {
		return appErr
	}

	data, appErr := users.Find(&id, nil)
	if appErr != nil {
		return appErr
	}

	record := data.Model

	//> determine which template to be provided
	isShow := true
	viewFile := "show"
	if strings.Contains(r.URL.Path, "/edit") {
		isShow = false
		// viewFile = "edit"
	}

	respData := map[string]interface{}{
		"title":          record.Name + "'s Profile",
		"previousUrl":    url.PreviousUrl(),
		"record":         record,
		"isShow":         isShow,
		csrf.TemplateTag: csrf.TemplateField(r),

		"success": ses.GetFlash("success"),
		"error":   ses.GetFlash("error"),
	}

	//> for api based
	if req.WantsJson() {
		return res.Json(respData, http.StatusOK)
	}

	//> for form based, show the "viewFile"
	return res.View(
		[]string{"base", fmt.Sprintf("users/%s", viewFile)},
		respData,
	)
}

func Update(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	message := "Successfully Updated!"
	status := http.StatusOK

	id := req.Input("id", nil).(string)
	data, appErr := users.Find(&id, nil)
	if appErr != nil {
		return appErr
	}
	data.Updates(req.All())

	//> for api based
	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
		}, status)
	}

	//> for form based, just redirect
	ses.SetFlash("success", message)
	url.RedirectPrevious()
	return nil
}