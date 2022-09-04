package lucid

import (
	"context"
	e "errors"
	"net/http"
	"time"

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

func (c Context) Engine() engines.EngineContract {
	return c.ctx.Value(EngineCtx{}).(engines.EngineContract)
}

func (c Context) Router() RouterContract {
	return c.ctx.Value(RouterCtx{}).(RouterContract)
}

func (c Context) Session() session.SessionContract {
	coo := c.Engine().GetCookie()
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

func (c Context) Stop() Middleware {
	return Middleware{
		Continue: false,
	}
}

func (c Context) Next() Middleware {
	return Middleware{
		Continue: true,
	}
}

func (c *Context) Bind(key interface{}, value interface{}) *Context {
	c.ctx = context.WithValue(c.ctx, key, value)
	return c
}

// this should be similar to Value()
func (c Context) Resolve(key any) any {
	return c.Value(key)
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c Context) Err() error {
	return c.ctx.Err()
}

func (c Context) Value(key any) any {
	return c.ctx.Value(key)
}
