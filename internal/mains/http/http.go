package main

import (
	"github.com/daison12006013/gorvel/internal/env"
	"github.com/daison12006013/gorvel/internal/kernel"
)

func main() {
	env.LoadEnv()
	kernel.HttpApplication()
}
