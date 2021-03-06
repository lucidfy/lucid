package registrar

import (
	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/app/commands"
	"github.com/lucidfy/lucid/pkg/facade/lang"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/loader"
	"github.com/lucidfy/lucid/pkg/lucid_commands"
	"github.com/lucidfy/lucid/resources/translations"
	cli "github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	commands.Inspire().Command,
}

// Lucid commands that needs to pass-in the template
// configurations such as Translations, Routes,
// GlobalMiddleware and Route middleware
var LucidCommands = []*cli.Command{
	lucid_commands.DatabaseMigration(Migrations),
	lucid_commands.RouteRegistered(func() *mux.Router {
		return loader.NetHttp(lang.Load(translations.Languages)).
			AddGlobalMiddlewares(GlobalMiddleware).
			AddRouteMiddlewares(RouteMiddleware).
			Register(&Routes)
	}),
	lucid_commands.RouteDefined(func() *[]routes.Routing {
		return loader.NetHttp(lang.Load(translations.Languages)).
			AddGlobalMiddlewares(GlobalMiddleware).
			AddRouteMiddlewares(RouteMiddleware).
			Explain(&Routes)
	}),
}
