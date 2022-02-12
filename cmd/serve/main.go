package main

import (
	"github.com/daison12006013/gorvel/internal/kernel"
	"github.com/daison12006013/gorvel/pkg/env"
)

func main() {
	env.LoadEnv()
	kernel.HttpApplication()
}
