package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	engine := engines.MuxEngine{Writer: w, Request: r}
	request := engine.ParsedRequest().(request.MuxRequest)
	response := engine.ParsedResponse().(response.MuxResponse)

	// prepare the data
	data := map[string]interface{}{
		"title": "Gorvel Rocks!",
	}

	// this is api request
	if request.IsJson() && request.WantsJson() {
		response.Json(data, http.StatusOK)
		return
	}

	// render the template
	response.View(
		// this example below, we're telling the compiler
		// to parse the base.go.html first, and then parse the welcome.go.html
		// therefore the defined "body" should render accordingly
		[]string{"base.go.html", "welcome.go.html"},
		data,
	)
}
