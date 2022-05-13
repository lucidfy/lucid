package auth_handler

import "github.com/lucidfy/lucid/pkg/facade/routes"

var RouteResource = routes.Routing{
	Path: "/auth/login",
	Name: "auth-login",
	Resources: routes.Resources{
		"index": User,         //  GET    /auth/login
		"store": LoginAttempt, //  POST   /auth/login
	},
}
