# Middleware

- [# Global Middleware](#-global-middleware)
- [# Route Middleware](#-route-middleware)
- [# Examples](#-examples)
  - [# Structure](#-structure)
  - [# Authenticated Middleware](#-authenticated-middleware)

---

A middleware intercepts the http request, you can throw an error, create a cookie, log the request or to validate any information.

> The configuration lives inside `app/kernel.go`

{#-global-middleware}

## [#](#-global-middleware) Global Middleware

The variable `GlobalMiddleware` every route will automatically process these middlewares, such as `middlewares.HttpAccessLogMiddleware`

{#-route-middleware}

## [#](#-route-middleware) Route Middleware

The variable `RouteMiddleware` will only be executed if you're going to specifically add it for each route, to learn more about this [Routing -> Middleware](/routing#-route-middlewares)

{#-examples}

## [#](#-examples) Examples

{#-structure}

### [#](#-structure) Structure

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // your conditions here

        // or else, process the next middleware
        next.ServeHTTP(w, r)
    })
}
```

{#-authenticated-middleware}

### [#](#-authenticated-middleware) Authenticated Middleware

Here's a sample middleware that lucid have, basically, we're checking if there was a `authenticated` inside the session.

If the session key isn't present, we then return a status `403 Forbidden`

```go
func AuthenticateMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ses := session.File(w, r)
        _, err := ses.Get("authenticated")

        if err != nil {
            handlers.HttpErrorHandler(engines.NetHttp(w, r), &errors.AppError{
                Code:    http.StatusForbidden,
                Message: "Forbidden!",
                Error:   err,
            })
            return
        }

        // if it passes, then we move on to the next middleware
        next.ServeHTTP(w, r)
    })
}
```

> To learn more about [Session](/session)
