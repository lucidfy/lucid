package registrar

import (
	"github.com/lucidfy/lucid/app/commands"
	"github.com/lucidfy/lucid/pkg/lucid_commands"
	cli "github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	commands.Inspire().Command,
	lucid_commands.RouteDefined(&Routes),
	lucid_commands.RouteRegistered(&Routes),
	lucid_commands.DatabaseMigration(Migrations),
}
