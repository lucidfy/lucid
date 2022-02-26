package kernel

import (
	"os"

	"github.com/daison12006013/gorvel/console"
	"github.com/daison12006013/gorvel/pkg/env"
	"github.com/daison12006013/gorvel/pkg/errors"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	app := &cli.App{
		Name:     "craft",
		Usage:    "A crafting console command tool for garvel!",
		Commands: *console.Commands,
	}

	err := app.Run(os.Args)
	if errors.Handler("error running craft console command", err) {
		panic(err)
	}
}
