package users_handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid"
)

func show(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	router := ctx.Router()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	bUrl, _ := router.Get("users.lists").URL()

	id := req.Input("id", nil).(string)
	if app_err := users.Exists("id", &id); app_err != nil {
		return app_err
	}

	data, app_err := users.Find(&id, nil)
	if app_err != nil {
		return app_err
	}

	record := data.Model

	//> determine which template to be provided
	is_show := true
	view_file := "show"
	if strings.Contains(url.CurrentURL(), "/edit") {
		is_show = false
		// view_file = "edit"
	}

	respData := map[string]interface{}{
		"title":          "Information",
		"record":         record,
		"isShow":         is_show,
		csrf.TemplateTag: csrf.TemplateField(engine.(engines.NetHttpEngine).HttpRequest),

		"success": ses.GetFlash("success"),
		"error":   ses.GetFlash("error"),

		"base_url": bUrl,
	}

	//> for api based
	if req.WantsJson() {
		return res.Json(respData, http.StatusOK)
	}

	//> for form based, show the "view_file"
	return res.View(
		[]string{"base", fmt.Sprintf("users/%s", view_file)},
		respData,
	)
}

func update(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	message := "Successfully Updated!"
	status := http.StatusOK

	id := req.Input("id", nil).(string)
	data, app_err := users.Find(&id, nil)
	if app_err != nil {
		return app_err
	}
	data.Updates(req.All())

	//> for api based
	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
		}, status)
	}

	//> for form based, just redirect
	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
