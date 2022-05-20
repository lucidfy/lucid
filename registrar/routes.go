package registrar

import (
	"github.com/lucidfy/lucid/app/handlers/auth_handler"
	"github.com/lucidfy/lucid/app/handlers/sample_handler"
	"github.com/lucidfy/lucid/app/handlers/single_handler"
	"github.com/lucidfy/lucid/app/handlers/users_handler"
	"github.com/lucidfy/lucid/pkg/facade/routes"
)

var Routes = &[]routes.Routing{
	{
		Path:   "/static",
		Name:   "static",
		Static: "./resources/static",
	},
	single_handler.WelcomeRoute,
	users_handler.RouteResource,
	auth_handler.RouteResource,
	sample_handler.RequestRoute,
	sample_handler.StorageRoute,
	sample_handler.DocsRoute,
}
