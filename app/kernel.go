package app

import (
	"github.com/daison12006013/gorvel/app/middlewares"
	"github.com/gorilla/mux"
)

// *-------------------------------------------------------------------
// * The application's global HTTP middleware stack.
// * These middleware are run during every request to your application.
// *-------------------------------------------------------------------
var Middleware = &[]mux.MiddlewareFunc{
	middlewares.HttpAccessLogMiddleware,
}

// *-------------------------------------------------------------------
// * The application's route middleware groups.
// *-------------------------------------------------------------------
var MiddlewareGroup = &map[string][]mux.MiddlewareFunc{
	"web": {
		// middlewares.EncryptCookiesMiddleware,
		// middlewares.AddQueuedCookiesToResponseMiddleware,
		// middlewares.StartSessionMiddleware,
		// // middlewares.AuthenticateSessionMiddleware,
		// middlewares.ShareErrorsFromSessionMiddleware,
		// middlewares.VerifyCsrfTokenMiddleware,
		// middlewares.SubstituteBindingsMiddleware,
	},
	"api": {
		// middlewares.ThrottleApiMiddleware,
		// middlewares.SubstituteBindingsMiddleware,
	},
}

// *-------------------------------------------------------------------
// * The application's route middleware.
// * These middleware may be assigned to groups or used individually.
// *-------------------------------------------------------------------
var RouteMiddleware = &map[string]mux.MiddlewareFunc{
	// "auth":             middlewares.AuthenticateMiddleware,
	// "auth.basic":       middlewares.AuthenticateWithBasicAuthMiddleware,
	// "cache.headers":    middlewares.SetCacheHeadersMiddleware,
	// "can":              middlewares.AuthorizeMiddleware,
	// "guest":            middlewares.RedirectIfAuthenticatedMiddleware,
	// "password.confirm": middlewares.RequirePasswordMiddleware,
	// "signed":           middlewares.ValidateSignatureMiddleware,
	// "throttle":         middlewares.ThrottleRequestsMiddleware,
	// "verified":         middlewares.EnsureEmailIsVerifiedMiddleware,
}
