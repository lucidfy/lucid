package usershandler

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
)

func Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logger.Info("Here at update!")
}
