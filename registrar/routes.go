package registrar

import (
	"github.com/lucidfy/lucid/pkg/facade/routes"
	"github.com/lucidfy/lucid/pkg/handlers/auth_handler"
	"github.com/lucidfy/lucid/pkg/handlers/sample_handler"
	"github.com/lucidfy/lucid/pkg/handlers/single_handler"
	"github.com/lucidfy/lucid/pkg/handlers/users_handler"
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
