package users_handler

import "github.com/lucidfy/lucid/pkg/facade/routes"

var RouteResource = routes.Routing{
	Path: "/users",
	Name: "users",
	Resources: routes.Resources{
		"index":   Lists,  //  GET    /users
		"create":  Create, //  GET    /users/create
		"store":   Store,  //  POST   /users
		"show":    Show,   //  GET    /users/{id}
		"edit":    Show,   //  GET    /users/{id}/edit
		"update":  Update, //  PUT    /users/{id}, POST /users/{id}/update
		"destroy": Delete, //  DELETE /users/{id}, POST /users/{id}/delete
	},
	// Middlewares: r.Middlewares{"auth"}, // Un-comment this line if you want to prevent it from guests
}
