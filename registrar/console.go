package registrar

import (
	"github.com/lucidfy/lucid/app/commands"
	cli "github.com/urfave/cli/v2"
)

var Commands = &[]*cli.Command{
	route_defined(),
	route_registered(),
	migrate(),
	commands.MakeInit().Command,
	commands.MakeHandler().Command,
	commands.MakeResource().Command,
	commands.MakeModel().Command,
	commands.MakeValidation().Command,
}

func route_defined() *cli.Command {
	return &cli.Command{
		Name:    "route:defined",
		Aliases: []string{"show:defined-routes"},
		Usage:   "Get the lists of defined routes",
		Action: func(c *cli.Context) error {
			commands.DefinedRoutes(c, Routes)
			return nil
		},
	}
}

func route_registered() *cli.Command {
	return &cli.Command{
		Name:    "route:registered",
		Aliases: []string{"show:routes"},
		Usage:   "Get the lists of registered routes",
		Action: func(c *cli.Context) error {
			commands.RegisteredRoutes(c, Routes)
			return nil
		},
	}
}

func migrate() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "To build your tables",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "current-database",
				Value: false,
				Usage: `Show the current database`,
			},
			&cli.BoolFlag{
				Name:  "auto-migrate",
				Value: false,
				Usage: `To automatically migrate the tables / columns`,
			},
			&cli.BoolFlag{
				Name:  "up",
				Value: false,
				Usage: `To increment migration`,
			},
			&cli.BoolFlag{
				Name:  "down",
				Value: false,
				Usage: `To rollback migration`,
			},
		},
		Action: func(c *cli.Context) error {
			commands.Migrate(c, Migrations)
			return nil
		},
	}
}
