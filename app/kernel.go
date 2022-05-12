package app

import (
	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/app/middlewares"
)

// Global Middleware
var Middleware = []mux.MiddlewareFunc{
	middlewares.HttpAccessLogMiddleware,
	middlewares.SessionPersistenceMiddleware,
	middlewares.CsrfShouldSkipMiddleware,
	middlewares.CsrfProtectMiddleware,
	middlewares.CsrfSetterMiddleware,
}

// Route Middleware
var RouteMiddleware = map[string]mux.MiddlewareFunc{
	"auth": middlewares.AuthenticateMiddleware,
}
