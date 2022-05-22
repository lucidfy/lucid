package lucid_commands

import (
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lucidfy/lucid/pkg/facade/routes"
	cli "github.com/urfave/cli/v2"
)

func RouteDefined(r *[]routes.Routing) *cli.Command {
	return &cli.Command{
		Name:    "route:defined",
		Aliases: []string{"show:rd"},
		Usage:   "Get the lists of defined routes",
		Action: func(c *cli.Context) error {
			defined := defined(r)

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Name", "Path", "Methods", "Middlewares", "Queries"})
			for _, routing := range defined {
				t.AppendRow(table.Row{routing.counter, routing.name, routing.path, routing.methods, routing.middlewares, routing.queries})
			}
			t.Render()
			return nil
		},
	}
}

func RouteRegistered(r *[]routes.Routing) *cli.Command {
	return &cli.Command{
		Name:    "route:registered",
		Aliases: []string{"show:rr"},
		Usage:   "Get the lists of registered routes",
		Action: func(c *cli.Context) error {

			registered := registered(r)

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Path", "Methods", "Path Regexp", "Queries", "Queries Regexp"})
			for _, routing := range registered {
				t.AppendRow(table.Row{routing.counter, routing.path, routing.methods, routing.pathregexp, routing.queries, routing.queriesregexp})
			}
			t.Render()
			return nil
		},
	}
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

func registered(r *[]routes.Routing) []Registered {
	routings := []Registered{}

	handlers := routes.NetHttp().Register(r)
	handlers.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}

		pathregexp, err := route.GetPathRegexp()
		if err != nil {
			pathregexp = ""
		}

		queries, err := route.GetQueriesTemplates()
		if err != nil {
			queries = []string{}
		}

		queriesregexp, err := route.GetQueriesRegexp()
		if err != nil {
			queries = []string{}
		}

		methods, err := route.GetMethods()
		if err != nil {
			methods = []string{"?"}
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

func defined(r *[]routes.Routing) []Defined {
	routings := []Defined{}

	routes := *routes.NetHttp().Explain(r)
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
