package users_handler

import "github.com/lucidfy/lucid/pkg/facade/routes"

var RouteResource = routes.Routing{
	Path: "/users",
	Name: "users",
	Resources: routes.Resources{
		"index":   lists,  //  GET    /users
		"create":  create, //  GET    /users/create
		"store":   store,  //  POST   /users
		"show":    show,   //  GET    /users/{id}
		"edit":    show,   //  GET    /users/{id}/edit
		"update":  update, //  PUT    /users/{id}, POST /users/{id}/update
		"destroy": delete, //  DELETE /users/{id}, POST /users/{id}/delete
	},
	// Middlewares: r.Middlewares{"auth"}, // Un-comment this line if you want to prevent it from guests
}
