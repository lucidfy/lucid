package usershandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/daison12006013/gorvel/app/models/users"
	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/gorilla/csrf"
)

func Show(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	request := engine.Request
	response := engine.Response

	id := request.Input("id", nil).(string)
	record := users.Find(&id).Record

	if request.IsJson() && request.WantsJson() {
		response.Json(record, http.StatusOK)
		return
	}

	// by default we use "show" then check if the
	// url path contains /edit , therefore use "edit"
	html := "show"
	if strings.Contains(engine.HttpRequest.URL.Path, "/edit") {
		html = "edit"
	}

	response.View(
		[]string{"base.go.html", fmt.Sprintf("users/%s.go.html", html)},
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

func Update(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	request := engine.Request
	response := engine.Response

	message := "Successfully Updated!"
	status := http.StatusOK

	id := request.Input("id", nil).(string)
	found := users.Find(&id)
	found.Updates(request.All())

	// for api based
	if request.IsJson() && request.WantsJson() {
		response.Json(map[string]interface{}{
			"ok":      true,
			"message": message,
		}, status)
		return
	}

	// for form based, just redirect
	request.SetFlash("success", message)
	request.RedirectPrevious()
}
