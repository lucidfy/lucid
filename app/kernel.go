package app

import (
	"github.com/daison12006013/gorvel/app/middlewares"
	"github.com/gorilla/mux"
)

// The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
var Middleware = []mux.MiddlewareFunc{
	middlewares.HttpAccessLogMiddleware,
	middlewares.SessionPersistenceMiddleware,
	middlewares.CsrfShouldSkipMiddleware,
	middlewares.CsrfProtectMiddleware,
	middlewares.CsrfSetterMiddleware,
}

// The application's route middleware.
// These middleware may be assigned to groups or used individually.
var RouteMiddleware = map[string]mux.MiddlewareFunc{
	"auth": middlewares.AuthenticateMiddleware,
}
