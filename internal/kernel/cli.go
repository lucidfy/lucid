package kernel

import (
	"fmt"
	"log"
	"os"

	"github.com/daison12006013/gorvel/internal/env"
	cli "github.com/urfave/cli/v2"
)

func ConsoleApplication() {
	env.LoadEnv()

	app := &cli.App{
		Name:  "craft",
		Usage: "A command to craft gorvel templates.",
		Commands: []*cli.Command{
			{
				Name:    "model",
				Aliases: []string{"m"},
				Usage:   "Creates a model",
				Action:  CraftModel,
			},
			{
				Name:    "handler",
				Aliases: []string{"h"},
				Usage:   "Creates an http handler",
				Action:  CraftHandler,
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func CraftModel(c *cli.Context) error {
	fmt.Println("make model ", c.Args().First())
	return nil
}

func CraftHandler(c *cli.Context) error {
	fmt.Println("make handler ", c.Args().First())
	return nil
}
