package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
)

func Create(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	// request := engine.Request
	response := engine.Response

	data := map[string]interface{}{"title": "Create Form"}

	response.View(
		[]string{"base.go.html", "users/create.go.html"},
		data,
	)
}

func Store(T engines.EngineInterface) {
	engine := T.(engines.MuxEngine)
	// request := engine.Request
	response := engine.Response

	// prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	// * TODO

	// prepare the data
	response.Json(map[string]interface{}{
		"message": message,
	}, status)
}
