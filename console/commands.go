package console

import "github.com/daison12006013/gorvel/app/commands"

func Init() Commands {
	return Commands{
		{
			Name:    "model",
			Aliases: Alias{"mo"},
			Usage:   "Creates a model",
			Action:  commands.CraftModel,
		},
		{
			Name:    "handler",
			Aliases: Alias{"hd"},
			Usage:   "Creates an http handler",
			Action:  commands.CraftHandler,
		},
		{
			Name:    "resource",
			Aliases: Alias{"res"},
			Usage:   "Creates an http resource handler",
			Action:  commands.CraftResource,
		},
		{
			Name:    "route:defined",
			Aliases: Alias{"ro:de"},
			Usage:   "Get the lists of defined routes",
			Action:  commands.DefinedRoutes,
		},
		{
			Name:    "route:registered",
			Aliases: Alias{"ro:re"},
			Usage:   "Get the lists of registered routes",
			Action:  commands.RegisteredRoutes,
		},
	}
}
