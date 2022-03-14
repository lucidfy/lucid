package helpers

import (
	"fmt"
	"log"
	"os"
)

type MP = map[string]interface{}

func DD(data ...interface{}) {
	prefix := fmt.Sprintf("[%s] [debug] (die-dump) -> ", os.Getenv("APP_ENV"))
	l := log.New(os.Stderr, prefix, log.LstdFlags)
	l.Printf("%+v\n", data...)
	os.Exit(1)
}
