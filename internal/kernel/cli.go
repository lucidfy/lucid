package kernel

import (
	"os"

	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid_commands"
	"github.com/lucidfy/lucid/registrar"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	cmds := lucid_commands.Commands
	cmds = append(cmds, registrar.Commands...)

	app := &cli.App{
		Name:     "Run",
		Usage:    "A console command runner for lucid!",
		Commands: cmds,
	}

	err := app.Run(os.Args)
	if errors.Handler("error running run console command", err) {
		panic(err)
	}
}
