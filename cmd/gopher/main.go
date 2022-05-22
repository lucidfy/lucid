package main

import (
	"github.com/lucidfy/lucid/internal/kernel"
	"github.com/lucidfy/lucid/registrar"
)

func main() {
	kernel.ConsoleApplication(registrar.Commands)
}
