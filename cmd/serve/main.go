package main

import (
	"flag"
	"os"

	"github.com/daison12006013/gorvel/internal/kernel"
	"github.com/daison12006013/gorvel/pkg/env"
)

func init() {
	env.LoadEnv()

	var host string
	var port string
	flag.StringVar(&host, "host", os.Getenv("HOST"), "Host to use")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "Port to use")
	flag.Parse()

	if len(host) > 0 {
		os.Setenv("HOST", host)
	}

	if len(port) > 0 {
		os.Setenv("PORT", port)
	}
}

func main() {
	kernel.
		Init().
		Run().
		WithGracefulShutdown()
}
