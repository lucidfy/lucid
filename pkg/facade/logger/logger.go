package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/daison12006013/gorvel/pkg/facade/path"
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

func New(prefix string) (*log.Logger, *os.File) {
	wrt, f := MakeWriter()
	logger := log.New(os.Stderr, prefix, log.LstdFlags)
	logger.SetOutput(wrt)
	return logger, f
}

func Info(title string, data ...interface{}) {
	logger, file := New(fmt.Sprintf("[%s] [info] ", os.Getenv("APP_ENV")))
	defer file.Close()
	logger.Printf("%s %v", title, data)
}

func Warning(title string, data ...interface{}) {
	logger, file := New(fmt.Sprintf("[%s] [warning] ", os.Getenv("APP_ENV")))
	defer file.Close()
	logger.Printf("%s %v", title, data)
}

func Error(title string, data ...interface{}) {
	logger, file := New(fmt.Sprintf("[%s] [error] ", os.Getenv("APP_ENV")))
	defer file.Close()
	logger.Printf("%s %v", title, data)
}

func Printf(format string, v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Printf(format, v...)
}

func Print(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Print(v...)
}

func Println(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Fatalln(v...)
}

func Panic(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	logger, file := New("")
	defer file.Close()
	logger.Panicln(v...)
}
