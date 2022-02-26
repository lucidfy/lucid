package commands

import (
	"os"
	"strings"

	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/routes"
	"github.com/daison12006013/gorvel/registrar"
	"github.com/gorilla/mux"
	"github.com/jedib0t/go-pretty/v6/table"
	cli "github.com/urfave/cli/v2"
)

func DefinedRoutes(c *cli.Context) error {
	defined := defined()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Path", "Methods", "Middlewares", "Queries"})
	for _, routing := range defined {
		t.AppendRow(table.Row{routing.counter, routing.name, routing.path, routing.methods, routing.middlewares, routing.queries})
	}
	t.Render()
	return nil
}

func RegisteredRoutes(c *cli.Context) error {
	registered := registered()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Path", "Methods", "Path Regexp", "Queries", "Queries Regexp"})
	for _, routing := range registered {
		t.AppendRow(table.Row{routing.counter, routing.path, routing.methods, routing.pathregexp, routing.queries, routing.queriesregexp})
	}
	t.Render()
	return nil
}

type Registered struct {
	counter       int
	path          string
	pathregexp    string
	queries       string
	queriesregexp string
	methods       string
}

type Defined struct {
	counter     int
	name        string
	path        string
	methods     string
	middlewares string
	queries     string
}

func registered() []Registered {
	routings := []Registered{}

	handlers := routes.Mux().Register(registrar.Routes).(*mux.Router)
	handlers.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		pathregexp, err := route.GetPathRegexp()
		if errors.Handler("GetPathRegexp", err) {
			return nil
		}

		queries, err := route.GetQueriesTemplates()
		if errors.Handler("GetQueriesTemplates", err) {
			return nil
		}

		queriesregexp, err := route.GetQueriesRegexp()
		if errors.Handler("GetQueriesRegexp", err) {
			return nil
		}

		methods, err := route.GetMethods()
		if errors.Handler("GetMethods", err) {
			return nil
		}

		routings = append(routings, Registered{
			counter:       len(routings) + 1,
			path:          path,
			pathregexp:    pathregexp,
			queries:       strings.Join(queries, ","),
			queriesregexp: strings.Join(queriesregexp, ","),
			methods:       strings.Join(methods, ","),
		})

		return nil
	})

	return routings
}

func defined() []Defined {
	routings := []Defined{}

	routes := *routes.Mux().Explain(registrar.Routes).(*[]routes.Routing)
	for _, route := range routes {
		routings = append(routings, Defined{
			counter:     len(routings) + 1,
			name:        route.Name,
			path:        route.Path,
			methods:     strings.Join(route.Method, ","),
			middlewares: strings.Join(route.Middlewares, ","),
			queries:     strings.Join(route.Queries, ","),
		})
	}
	return routings
}
