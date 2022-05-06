# Routing

Lucid's routing structure is simple, as if you're just writing a json schema.

## Structure

Routes are stored inside `registrar/routes.go` under a variable called `Routes`

Here is a simple way to write a route.

```go
var Routes = &[]r.Routing{
    {
		Path:    "/",
		Name:    "welcome",
		Method:  r.Method{"GET"},
		Handler: handlers.Welcome,
    },
}
```

To explain above

The `Path` is your **url** relative path, while the `Handler` shall be the one to literally handle the route, if you're coming from different frameworks, the alternative term for this is `Controller` or `Action`.

> `Tips:` some engineers tend to have a `Service` or `Repository` pattern, you can apply the same way but most of the time you don't need to over do things, although it is good to extract your logic into pieces to easily apply a unit test.
> To learn more about [Handlers](/handlers)

Meanwhile, the `Method` is the action that is coming from your browser's form, XMLHttpRequest (ajax / fetch)

HTTP Method | CRUD Operation
---------|----------
 GET | Read
 POST | Create
 PATCH | Update
 DELETE | Delete
 PUT | Update/Replace

> `Note:` if you want to handle multiple method in one route, you just need to append the value using a comma separated value.

```go
r.Method{"GET", "POST"},
```

## Route Resource

We're commonly building routes to serve a [C.R.U.D.](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete), such as **Users Management**, **Orders**, **Reports** and more!

```go
...
{
    Path: "/users",
    Name: "users",
    Resources: r.Resources{
        "index":   usershandler.Lists,
        "create":  usershandler.Create,
        "store":   usershandler.Store,
        "show":    usershandler.Show,
        "edit":    usershandler.Show,
        "update":  usershandler.Update,
        "destroy": usershandler.Delete,
    },
},
```

Here's how lucid will understand the routing resource

Default |  Path | Alternative Path
---------|----------|----------
index | `GET` /users | -
create | `GET` /users/create | -
store | `POST` /users | -
show | `GET` /users/{id} | -
edit | `GET` /users/{id}/edit | -
update | `PUT` /users/{id} | `POST` /users/{id}/update
destroy | `DELETE` /users/{id} | `POST` /users/{id}/delete

To learn more about the core behind this, please read [Core -> Routing Resource](/core-routing-resource)

## Route Middlewares

Middleware is used to intercept the url request before it goes to our handler

```go
{
    Path: "/users",
    ...
    Middlewares: r.Middlewares{"auth"},
}
```

As an example above, we're injecting the `auth`, this string is stored inside `app/kernel.go` and hooked under `middlewares.AuthenticateMiddleware`

```go
var RouteMiddleware = map[string]mux.MiddlewareFunc{
	"auth": middlewares.AuthenticateMiddleware,
}
```

> `Note:` for more info on how a middleware works [Middlewares](/middleware)
