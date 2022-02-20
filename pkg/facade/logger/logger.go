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
	l := log.New(os.Stderr, prefix, log.LstdFlags)
	l.SetOutput(wrt)
	return l, f
}

func Debug(title string, data ...interface{}) {
	text("[%s] [debug] ", title, data...)
}

func Info(title string, data ...interface{}) {
	text("[%s] [info] ", title, data...)
}

func Warning(title string, data ...interface{}) {
	text("[%s] [warning] ", title, data...)
}

func Error(title string, data ...interface{}) {
	text("[%s] [error] ", title, data...)
}

func text(txt string, title string, data ...interface{}) {
	l, file := New(fmt.Sprintf(txt, os.Getenv("APP_ENV")))
	defer file.Close()
	data = prepend(title, data...)
	l.Printf("%s %+v\n", data...)
}

func prepend(addtl interface{}, data ...interface{}) []interface{} {
	data = append(data, 0)
	copy(data[1:], data)
	data[0] = addtl
	return data
}

func Printf(format string, v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Printf(format, v...)
}

func Print(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Print(v...)
}

func Println(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Println(v...)
}

func Fatal(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Fatalln(v...)
}

func Panic(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	l, file := New("")
	defer file.Close()
	l.Panicln(v...)
}
