package main

import (
	"github.com/lucidfy/lucid/internal/kernel"
	"github.com/lucidfy/lucid/registrar"
)

func main() {
	cmds := registrar.Commands
	cmds = append(cmds, registrar.LucidCommands...)
	kernel.ConsoleApplication(cmds...)
}
