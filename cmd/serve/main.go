package main

import (
	"github.com/daison12006013/gorvel/internal/kernel"
	"github.com/daison12006013/gorvel/pkg/env"
)

func init() {

}

func main() {
	env.LoadEnv()
	kernel.
		Init().
		Run().
		WithGracefulShutdown()
}
