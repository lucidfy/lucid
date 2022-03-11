package registrar

import (
	"github.com/daison12006013/gorvel/app/handlers"
	"github.com/daison12006013/gorvel/app/handlers/usershandler"
	r "github.com/daison12006013/gorvel/pkg/facade/routes"
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
	// {
	// 	Path:     "/docs",
	// 	Name:     "docs",
	// 	Markdown: "resources/docs",
	// },
	{
		Path:   "/static",
		Name:   "static",
		Static: "./resources/static",
	},
}
