package lucid

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/engines"
)

type EngineCtx struct{}
type RouterCtx struct{}

type Context struct {
	ctx context.Context
}

func NewContext(ctx context.Context) Context {
	return Context{
		ctx: ctx,
	}
}

func (resolver Context) Engine() engines.EngineContract {
	return resolver.ctx.Value(EngineCtx{}).(engines.EngineContract)
}

func (resolver Context) Router() *mux.Router {
	return resolver.ctx.Value(RouterCtx{}).(*(mux.Router))
}
