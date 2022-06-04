package routes

import (
	"context"

	"github.com/lucidfy/lucid/pkg/errors"
)

// structs

type Middlewares []string
type Queries []string
type Method []string
type Resources map[string]Handler
type Handler func(context.Context) *errors.AppError // func(engines.EngineContract) *errors.AppError
type Routing struct {
	Name                 string
	Path                 string
	Prefix               bool
	Method               []string
	Queries              Queries
	Handler              Handler
	Resources            map[string]Handler
	Middlewares          []string
	Static               string
	WithGlobalMiddleware interface{}
}
