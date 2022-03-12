package middlewares

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
)

func IsJsonRequest(w http.ResponseWriter, r *http.Request) bool {
	rp := request.Mux(w, r)

	if rp.IsJson() && rp.WantsJson() {
		return true
	}

	return false
}
