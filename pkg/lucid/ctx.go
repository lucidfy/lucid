package lucid

import (
	"context"
	e "errors"
	"net/http"

	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/facade/session"
	"github.com/lucidfy/lucid/pkg/helpers"
)

type EngineCtx struct{}
type RouterCtx struct{}
type Middleware struct {
	Continue bool
}
type Context struct {
	ctx context.Context
}

func New(ctx context.Context) Context {
	return Context{ctx: ctx}
}

func (resolver Context) Original() context.Context {
	return resolver.ctx
}

func (resolver Context) Engine() engines.EngineContract {
	return resolver.ctx.Value(EngineCtx{}).(engines.EngineContract)
}

func (resolver Context) Router() RouterContract {
	return resolver.ctx.Value(RouterCtx{}).(RouterContract)
}

func (resolver Context) Session() session.SessionContract {
	coo := resolver.Engine().GetCookie()
	var ses session.SessionContract

	if helpers.IsTrue(helpers.Getenv("SESSION_ENABLED", "true")) {
		sessionKey, app_err := coo.Get(helpers.Getenv("SESSION_NAME", "lucid_session"))
		if app_err != nil && e.Is(app_err.Error, http.ErrNoCookie) {
			sessionKey = coo.CreateSessionCookie()
		}

		ses = session.Driver(
			helpers.Getenv("SESSION_DRIVER", "file"),
			sessionKey.(string),
		)
	}

	return ses
}

func (resolver Context) Stop() Middleware {
	return Middleware{
		Continue: false,
	}
}

func (resolver Context) Next() Middleware {
	return Middleware{
		Continue: true,
	}
}
