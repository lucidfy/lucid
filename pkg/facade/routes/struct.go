package routes

import (
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid"
)

// structs

type Middlewares []string
type Queries []string
type Method []string
type Resources map[string]Handler
type Handler func(lucid.Context) *errors.AppError
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
