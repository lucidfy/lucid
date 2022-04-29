package registrar

import (
	"github.com/daison12006013/lucid/app/commands"
	cli "github.com/urfave/cli/v2"
)

var Commands = &[]*cli.Command{
	{
		Name:    "make:model",
		Aliases: []string{"m:m"},
		Usage:   "Creates a model",
		Action:  commands.CraftModel,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The name of your package (i.e: "profiles")`,
			},
			&cli.StringFlag{
				Name:  "table",
				Value: "",
				Usage: `The name of your table (i.e: "user_profiles")`,
			},
		},
	},
	{
		Name:    "make:handler",
		Aliases: []string{"m:h"},
		Usage:   "Creates an http handler",
		Action:  commands.CraftHandler,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The name of your handler (i.e: "users")`,
			},
		},
	},
	{
		Name:    "make:resource",
		Aliases: []string{"m:r"},
		Usage:   "Creates a resource along with the model",
		Action:  commands.CraftResource,
	},
	{
		Name:    "route:defined",
		Aliases: []string{"ro:de"},
		Usage:   "Get the lists of defined routes",
		Action: func(c *cli.Context) error {
			commands.DefinedRoutes(c, Routes)
			return nil
		},
	},
	{
		Name:    "route:registered",
		Aliases: []string{"ro:re"},
		Usage:   "Get the lists of registered routes",
		Action: func(c *cli.Context) error {
			commands.RegisteredRoutes(c, Routes)
			return nil
		},
	},
}
