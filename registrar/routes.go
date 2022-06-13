package registrar

import (
	"github.com/lucidfy/lucid/app/handlers/auth_handler"
	"github.com/lucidfy/lucid/app/handlers/sample_handler"
	"github.com/lucidfy/lucid/app/handlers/single_handler"
	"github.com/lucidfy/lucid/app/handlers/users_handler"
	"github.com/lucidfy/lucid/pkg/facade/routes"
)

var Routes = []routes.Routing{
	single_handler.WelcomeRoute,
	single_handler.StaticFiles,
	sample_handler.DocsRoute,
	auth_handler.RouteResource,
	users_handler.RouteResource,
	sample_handler.RequestRoute,
	sample_handler.StorageRoute,
}
