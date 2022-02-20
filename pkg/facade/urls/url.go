package urls

import (
	"os"
)

func GetAddr() string {
	var port string
	if len(os.Getenv("PORT")) > 0 {
		port = ":" + os.Getenv("PORT")
	}
	return os.Getenv("HOST") + port
}

func BaseUrl(uri *string) string {
	return os.Getenv("SCHEMA") + "://" + GetAddr() + *uri
}