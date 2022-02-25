package usershandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/daison12006013/gorvel/pkg/engines"
	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Show(w http.ResponseWriter, r *http.Request) {
	engine := engines.MuxEngine{Writer: w, Request: r}
	request := engine.ParsedRequest().(request.MuxRequest)
	response := engine.ParsedResponse().(response.MuxResponse)

	// // fetch the record in the database
	// record, err := users.FindById(*req.Input("id"))
	// if err != nil {
	// 	// if we're on debugging mode, just throw the error
	// 	if os.Getenv("APP_DEBUG") == "true" {
	// 		logger.Fatal(err)
	// 	}
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	// // prepare the data
	data := map[string]interface{}{
		"previousUrl": request.PreviousUrl(),
	}
	// data := map[string]interface{}{
	// 	"title":  record.Name + "'s Profile",
	// 	"record": record,
	// }

	// // this is api request
	// if req.IsJson() && req.WantsJson() {
	// 	response.Json(w, data, http.StatusOK)
	// 	return
	// }

	// by default we use "show"
	// then check if the url path contains /edit
	// therefore use "edit"
	html := "show"
	if strings.Contains(r.URL.Path, "/edit") {
		html = "edit"
	}

	response.View(
		[]string{"base.go.html", fmt.Sprintf("users/%s.go.html", html)},
		data,
	)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
