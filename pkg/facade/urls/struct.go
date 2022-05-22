package urls

import (
	"os"
)

type URLContract interface {
	CurrentURL() string
	FullURL() string
	PreviousURL() string
	RedirectPrevious()
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
