package middlewares

import (
	"net/http"

	"github.com/lucidfy/lucid/pkg/facade/request"
)

// IsJsonRequest check if request is json
// Content-Type: application/json
// Accept: application/json
func IsJsonRequest(w http.ResponseWriter, r *http.Request) bool {
	rp := request.Mux(w, r, nil)

	if rp.IsJson() || rp.WantsJson() {
		return true
	}

	return false
}
