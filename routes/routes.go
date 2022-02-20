package routes

import (
	"github.com/daison12006013/gorvel/app/handlers"
	"github.com/daison12006013/gorvel/app/handlers/usershandler"
)

func Routes() *[]Routing {
	r := &[]Routing{
		{
			Path:    "/",
			Name:    "welcome",
			Method:  Method{"GET"},
			Handler: handlers.Home,
		},
		{
			Path: "/users",
			Name: "users",
			Resources: Resources{
				"index":   usershandler.Lists,  //  GET    /users
				"create":  usershandler.Create, //  GET    /users/create
				"store":   usershandler.Store,  //  POST   /users
				"show":    usershandler.Show,   //  GET    /users/{id}
				"edit":    usershandler.Show,   //  GET    /users/{id}/edit
				"update":  usershandler.Update, //  PUT    /users/{id}, POST /users/{id}/update
				"destroy": usershandler.Delete, //  DELETE /users/{id}, POST /users/{id}/delete
			},
			Middlewares: Middlewares{"auth"},
		},
	}
	return r
}
