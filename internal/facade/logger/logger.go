package logger

import (
	"io"
	"log"
	"os"

	"github.com/daison12006013/gorvel/internal/facade/path"
)

func MakeWriter() (io.Writer, *os.File) {
	f, err := os.OpenFile(
		path.PathTo(os.Getenv("LOG_FILE")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666,
	)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, f)

	return wrt, f
}

func Info(title string, data ...interface{}) {
	wrt, f := MakeWriter()
	defer f.Close()

	log.SetOutput(wrt)
	log.Printf("[info] %s: %v", title, data)
}

func Error(title string, data ...interface{}) {
	wrt, f := MakeWriter()
	defer f.Close()

	log.SetOutput(wrt)
	log.Printf("[error] %s: %v", title, data)
}
