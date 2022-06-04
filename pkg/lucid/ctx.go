package lucid

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/lucidfy/lucid/pkg/engines"
)

type EngineCtx struct{}
type RouterCtx struct{}

type ContextResolver struct {
	ctx context.Context
}

func Context(ctx context.Context) *ContextResolver {
	return &ContextResolver{
		ctx: ctx,
	}
}

func (resolver ContextResolver) Engine() engines.EngineContract {
	return resolver.ctx.Value(EngineCtx{}).(engines.EngineContract)
}

func (resolver ContextResolver) Router() *mux.Router {
	return resolver.ctx.Value(RouterCtx{}).(*(mux.Router))
}
