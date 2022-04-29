package app

import (
	"github.com/daison12006013/lucid/app/middlewares"
	"github.com/gorilla/mux"
)

// Middleware The application's global HTTP middleware stack.
// This middlewares are run during every request to your application.
var Middleware = []mux.MiddlewareFunc{
	middlewares.HttpAccessLogMiddleware,
	middlewares.SessionPersistenceMiddleware,
	middlewares.CsrfShouldSkipMiddleware,
	middlewares.CsrfProtectMiddleware,
	middlewares.CsrfSetterMiddleware,
}

// RouteMiddleware The application's route middleware.
// These middlewares may be assigned to group's or used individually.
var RouteMiddleware = map[string]mux.MiddlewareFunc{
	"auth": middlewares.AuthenticateMiddleware,
}
