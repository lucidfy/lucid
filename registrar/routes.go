package registrar

import (
	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/app/handlers/authhandler"
	"github.com/lucidfy/lucid/app/handlers/samplehandler"
	"github.com/lucidfy/lucid/app/handlers/usershandler"
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
			"index":   usershandler.Lists,  //  GET    /users
			"create":  usershandler.Create, //  GET    /users/create
			"store":   usershandler.Store,  //  POST   /users
			"show":    usershandler.Show,   //  GET    /users/{id}
			"edit":    usershandler.Show,   //  GET    /users/{id}/edit
			"update":  usershandler.Update, //  PUT    /users/{id}, POST /users/{id}/update
			"destroy": usershandler.Delete, //  DELETE /users/{id}, POST /users/{id}/delete
		},
		Middlewares: r.Middlewares{"auth"},
	},
	{
		Path:    "/samples/requests",
		Name:    "",
		Method:  r.Method{"GET", "POST"},
		Handler: samplehandler.Requests,
	},
	{
		Path:    "/samples/storage",
		Name:    "",
		Method:  r.Method{"POST"},
		Handler: samplehandler.FileStorage,
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
			"index": authhandler.User,
			"store": authhandler.LoginAttempt,
		},
	},
}
