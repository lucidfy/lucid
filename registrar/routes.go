package registrar

import (
	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/app/handlers/auth_handler"
	"github.com/lucidfy/lucid/app/handlers/sample_handler"
	"github.com/lucidfy/lucid/app/handlers/users_handler"
	r "github.com/lucidfy/lucid/pkg/facade/routes"
)

var Routes = &[]r.Routing{
	{
		Path:    "/",
		Name:    "welcome",
		Method:  r.Method{"GET"},
		Handler: handlers.Welcome,
	},
	{
		Path: "/users",
		Name: "users",
		Resources: r.Resources{
			"index":   users_handler.Lists,  //  GET    /users
			"create":  users_handler.Create, //  GET    /users/create
			"store":   users_handler.Store,  //  POST   /users
			"show":    users_handler.Show,   //  GET    /users/{id}
			"edit":    users_handler.Show,   //  GET    /users/{id}/edit
			"update":  users_handler.Update, //  PUT    /users/{id}, POST /users/{id}/update
			"destroy": users_handler.Delete, //  DELETE /users/{id}, POST /users/{id}/delete
		},
		Middlewares: r.Middlewares{"auth"},
	},
	{
		Path:    "/samples/requests",
		Name:    "",
		Method:  r.Method{"GET", "POST"},
		Handler: sample_handler.Requests,
	},
	{
		Path:    "/samples/storage",
		Name:    "",
		Method:  r.Method{"POST"},
		Handler: sample_handler.FileStorage,
	},
	{
		Path:   "/static",
		Name:   "static",
		Static: "./resources/static",
	},
	{
		Path:    "/docs",
		Prefix:  true,
		Name:    "docs",
		Method:  r.Method{"GET"},
		Handler: handlers.Docs,
	},
	{
		Path: "/auth/login",
		Name: "auth-login",
		Resources: r.Resources{
			"index": auth_handler.User,
			"store": auth_handler.LoginAttempt,
		},
	},
}
