package helpers

import (
	"os"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
)

func DD(data ...interface{}) {
	logger.Debug("(die-dump) ->", data...)
	os.Exit(1)
}
