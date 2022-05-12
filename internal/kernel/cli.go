package kernel

import (
	"os"

	"github.com/lucidfy/lucid/pkg/env"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/registrar"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	app := &cli.App{
		Name:     "Run",
		Usage:    "A console command runner for lucid!",
		Commands: *registrar.Commands,
	}

	err := app.Run(os.Args)
	if errors.Handler("error running run console command", err) {
		panic(err)
	}
}
