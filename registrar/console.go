package registrar

import (
	"github.com/lucidfy/lucid/app/commands"
	cli "github.com/urfave/cli/v2"
)

var Commands = &[]*cli.Command{
	{
		Name:    "route:defined",
		Aliases: []string{"show:defined-routes"},
		Usage:   "Get the lists of defined routes",
		Action: func(c *cli.Context) error {
			commands.DefinedRoutes(c, Routes)
			return nil
		},
	},
	{
		Name:    "route:registered",
		Aliases: []string{"show:routes"},
		Usage:   "Get the lists of registered routes",
		Action: func(c *cli.Context) error {
			commands.RegisteredRoutes(c, Routes)
			return nil
		},
	},
	commands.CraftModel().Command,
	commands.CraftHandler().Command,
	commands.CraftResource().Command,
	commands.CraftValidation().Command,
}
