package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Create(w http.ResponseWriter, r *http.Request) {
	engine := engines.MuxEngine{Writer: w, Request: r}
	// request := engine.ParsedRequest().(request.MuxRequest)
	response := engine.ParsedResponse().(response.MuxResponse)

	data := map[string]interface{}{"title": "Create Form"}

	response.View(
		[]string{"base.go.html", "users/create.go.html"},
		data,
	)
}

func Store(w http.ResponseWriter, r *http.Request) {
	engine := engines.MuxEngine{Writer: w, Request: r}
	// request := engine.ParsedRequest().(request.MuxRequest)
	response := engine.ParsedResponse().(response.MuxResponse)

	// prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	// * TODO

	// prepare the data
	response.Json(map[string]interface{}{
		"message": message,
	}, status)
}
