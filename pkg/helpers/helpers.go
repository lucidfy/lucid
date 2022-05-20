package helpers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lucidfy/lucid/pkg/errors"
)

type MP = map[string]interface{}
type MS = map[string]string

func DD(data ...interface{}) {
	prefix := fmt.Sprintf("[%s] [debug] (die-dump) -> ", os.Getenv("APP_ENV"))
	l := log.New(os.Stderr, prefix, log.LstdFlags)
	l.Printf("%+v\n", data...)
	os.Exit(1)
}

func StringToInt(s string) (i int, app_err *errors.AppError) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1, &errors.AppError{
			Message: "helpers.StringToInt() parsing error",
			Code:    http.StatusInternalServerError,
			Error:   err,
		}
	}
	return i, nil
}
