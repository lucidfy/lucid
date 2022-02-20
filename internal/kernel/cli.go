package kernel

import (
	"os"

	"github.com/daison12006013/gorvel/console"
	"github.com/daison12006013/gorvel/pkg/env"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	app := &cli.App{
		Name:     "craft",
		Usage:    "A command to craft gorvel templates.",
		Commands: console.Init(),
	}

	err := app.Run(os.Args)

	if err != nil {
		logger.Fatal(err)
	}
}
