package middlewares

import (
	e "errors"
	"net/http"

	"github.com/lucidfy/lucid/app/handlers"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid"
)

func AuthenticateMiddleware(ctx lucid.Context) lucid.Middleware {
	engine := ctx.Engine()
	ses := ctx.Session()

	if ses == nil {
		handlers.HttpErrorHandler(engine, &errors.AppError{
			Code:    http.StatusForbidden,
			Message: "Session not found!",
			Error:   e.New("session not found"),
		}, nil)
		return ctx.Stop()
	}

	authorized, app_err := ses.Get("authenticated")
	if authorized == nil {
		handlers.HttpErrorHandler(engine, &errors.AppError{
			Code:    http.StatusForbidden,
			Message: "Forbidden!",
			Error:   e.New("you are not authorized"),
		}, nil)
		return ctx.Stop()
	}

	if app_err != nil {
		handlers.HttpErrorHandler(engine, &errors.AppError{
			Code:    http.StatusForbidden,
			Message: "Forbidden!",
			Error:   app_err.Error,
		}, nil)
		return ctx.Stop()
	}

	return ctx.Next()
}
