package errors

import (
	"net/http"
	"os"

	"github.com/lucidfy/lucid/pkg/facade/logger"
)

type AppError struct {
	Error           error
	Message         interface{}
	Code            interface{}
	ValidationError interface{}
}

func Handler(title string, e error) bool {
	if e != nil {
		// if we're on debugging mode
		// log the error
		if os.Getenv("APP_DEBUG") == "true" {
			logger.Error(title, e)
		}
		return true
	}
	return false
}

func InternalServerError(message string, err error) *AppError {
	if err != nil {
		return &AppError{
			Message: message,
			Code:    http.StatusInternalServerError,
			Error:   err,
		}
	}
	return nil
}
