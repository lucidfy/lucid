package urls

import (
	"net/url"
	"os"

	"github.com/lucidfy/lucid/pkg/errors"
)

type URLContract interface {
	Default() *url.URL

	BaseURL() string
	CurrentURL() string
	PreviousURL() string
	RedirectPrevious() *errors.AppError
}

func GetAddr() string {
	var port string
	if len(os.Getenv("PORT")) > 0 {
		port = ":" + os.Getenv("PORT")
	}
	return os.Getenv("HOST") + port
}

func BaseURL(uri *string) string {
	var u string = ""
	if uri != nil {
		u = *uri
	}
	return os.Getenv("SCHEME") + "://" + GetAddr() + u
}
