package console

import (
	"github.com/daison12006013/gorvel/app/commands"
	cli "github.com/urfave/cli/v2"
)

var Commands = &[]*cli.Command{
	{
		Name:    "model",
		Aliases: []string{"mo"},
		Usage:   "Creates a model",
		Action:  commands.CraftModel,
	},
	{
		Name:    "handler",
		Aliases: []string{"hd"},
		Usage:   "Creates an http handler",
		Action:  commands.CraftHandler,
	},
	{
		Name:    "resource",
		Aliases: []string{"res"},
		Usage:   "Creates a resource along with the model",
		Action:  commands.CraftResource,
	},
	{
		Name:    "route:defined",
		Aliases: []string{"ro:de"},
		Usage:   "Get the lists of defined routes",
		Action:  commands.DefinedRoutes,
	},
	{
		Name:    "route:registered",
		Aliases: []string{"ro:re"},
		Usage:   "Get the lists of registered routes",
		Action:  commands.RegisteredRoutes,
	},
}
