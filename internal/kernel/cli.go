package kernel

import (
	"os"

	"github.com/daison12006013/lucid/pkg/env"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/registrar"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	app := &cli.App{
		Name:     "craft",
		Usage:    "A crafting console command tool for garvel!",
		Commands: *registrar.Commands,
	}

	err := app.Run(os.Args)
	if errors.Handler("error running craft console command", err) {
		panic(err)
	}
}
