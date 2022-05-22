package kernel

import (
	"os"

	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid_commands"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication(cmds []*cli.Command) {
	env.LoadEnv()
	cmds = append(cmds, lucid_commands.Commands...)

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
