package app

import (
	"github.com/lucidfy/lucid/app/middlewares"
)

// Global Middleware
var GlobalMiddleware = []interface{}{
	middlewares.HttpAccessLogMiddleware,
	middlewares.SessionPersistenceMiddleware,
	middlewares.CsrfShouldSkipMiddleware,
	middlewares.CsrfProtectMiddleware,
	middlewares.CsrfSetterMiddleware,
}

// Route Middleware
var RouteMiddleware = map[string]interface{}{
	"auth": middlewares.AuthenticateMiddleware,
}
