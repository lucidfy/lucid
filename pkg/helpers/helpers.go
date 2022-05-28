package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

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

func Getenv(key string, d string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return d
}

func Stringify(data interface{}) string {
	var content string
	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		j, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		content = string(j)
	case reflect.String:
		content = data.(string)
	}

	return content
}

func SessionName() string {
	return Getenv("SESSION_NAME", "lucid_session")
}

func IsTrue(s string) bool {
	s = strings.ToLower(s)
	if s == "true" || s == "1" || s == "yes" {
		return true
	}
	return false
}
