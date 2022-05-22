# Handlers

- [# Basics](#-basics)
  - [# Engine](#-engine)
  - [# Request & Response](#-request--response)
  - [# Welcome Page](#-welcome-page)
- [# Console Command](#-console-command)
  - [# Single Handler](#-single-handler)
  - [# Resource Handler](#-resource-handler)

Handlers are the one responding to any http requests, the function will only be accessible after the middlewares had been iterated.

---

{#-basics}

## [#](#-basics) Basics

{#-engine}

### [#](#-engine) Engine

```go
func Sample(T engines.EngineContract) *errors.AppError {
    engine := T.(engines.NetHttpEngine)
}
```

As of writing, we're currently using [gorilla/mux](https://github.com/gorilla/mux) to bootstrap our routing and handlers, although we customized it to be called as "Engine".

> An analogy for this, an Engine differs depending on consumer needs, for instance, a consumer wanted to have a Diesel Engine for speed and cost-effectiveness although there are drawbacks about it, similarly if a consumer wanted to have a Petrol Engine.

> This is the same way for us [Software Engineers](https://en.wikipedia.org/wiki/Software_engineering), we can use an engine for you to replace it anytime you want, although as of the moment, we're expanding the lists of Engines to be supported, or by the community itself.

> Future plans, we're thinking to have this [Fiber](https://github.com/gofiber/fiber) as the speed is really promising, by transitioning to this engine sooner, it will just be easy for us!

{#-request--response}

### [#](#-request--response) Request & Response

```go
engine := T.(engines.NetHttpEngine)
w := engine.HttpResponseWriter
r := engine.HttpRequest
request := engine.Request
response := engine.Response
```

The variable `w` and `r` were based from [net/http](https://pkg.go.dev/net/http), they are commonly known as `http.ResponseWriter` and `*http.Request` in the go community.

While the `request` and `response` were interpreters to some of the top framework nowadays, lucid is heavely inspired by Laravel so these 2 variables contains at leasts the common functions such `Input()` / `WantsJson()`, and so on and so fort.

> To learn more, you can check out the `request` functions [here](/api-request), while `response` functions can be seen [here](/api-response)

{#-welcome-page}

### [#](#-welcome-page) Welcome Page

Here's a full example of our welcome page, the explaination can be seen inside as comment block

```go
package handlers

import (
    "net/http"
    "github.com/lucidfy/lucid/pkg/engines"
    "github.com/lucidfy/lucid/pkg/errors"
)

func Welcome(T engines.EngineContract) *errors.AppError {
    engine := T.(engines.NetHttpEngine)
    request := engine.Request
    response := engine.Response

    // let us initialize how to create a data
    data := map[string]interface{}{
        "title": "Lucid Rocks!",
    }

    // here, we're checking if the request wants a json response
    // we then pass the above data as a json, with StatusOK (200)
    if request.WantsJson() {
        return response.Json(data, http.StatusOK)
    }

    // go has its own template engine, to interpret below:
    // -> it parses the base.go.html (provided with the map string `data`)
    // -> after that it parses the welcome.go.html (provided with the map string `data`)
    return response.View(
        []string{"base", "welcome"},
        data,
    )
}
```


{#-console-command}

## [#](#-console-command) Console Command

These console commands will help you to generate a `go` file to speed up your development!

{#-single-handler}

### [#](#-single-handler) Single Handler

Here's how to generate a single handler

```bash
./run make:handler health_check

Created handler, located at:
 > ~/lucid-setup/src/lucid/app/handlers/health_check.go

Go to registrar/routes.go and paste this:

    var Routes = &[]routes.Routing{
        ...,
        {
            Path:    "/health-check",
            Name:    "health-check",
            Method:  routes.Method{"GET"}, // defaulting to "GET"
            Handler: handlers.HealthCheck,
        },
    }
```

{#-resource-handler}

### [#](#-resource-handler) Resource Handler

To generate a resource handler

```bash
./run make:resource reports

Created model, located at:
 > ~/lucid-setup/src/lucid/app/models/reports/struct.go
 > ~/lucid-setup/src/lucid/app/models/reports/model_test.go
 > ~/lucid-setup/src/lucid/app/models/reports/model.go

Created validation, located at:
 > ~/lucid-setup/src/lucid//app/validations/reports.go

Created resource handler, located at:
 > ~/lucid-setup/src/lucid/app/handlers/reports_handler/delete.go
 > ~/lucid-setup/src/lucid/app/handlers/reports_handler/lists.go
 > ~/lucid-setup/src/lucid/app/handlers/reports_handler/update.go
 > ~/lucid-setup/src/lucid/app/handlers/reports_handler/struct.go
 > ~/lucid-setup/src/lucid/app/handlers/reports_handler/create.go

Go to registrar/routes.go and paste this:

    var Routes = &[]routes.Routing{
        ...
        reports_handler.RouteResource,
    }
```
