package single_handler

import "github.com/lucidfy/lucid/pkg/facade/routes"

var StaticFiles = routes.Routing{
	Path:   "/static",
	Name:   "static",
	Static: "./resources/static",
}
