package auth_handler

import "github.com/lucidfy/lucid/pkg/facade/routes"

var RouteResource = routes.Routing{
	Path: "/auth/login",
	Name: "auth-login",
	Resources: routes.Resources{
		"index": user,          //  GET    /auth/login
		"store": login_attempt, //  POST   /auth/login
	},
}
