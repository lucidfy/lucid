package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	cli "github.com/urfave/cli/v2"
)

type MakeHandlerCommand struct {
	Command *cli.Command
}

func MakeHandler() *MakeHandlerCommand {
	var cc MakeHandlerCommand
	cc.Command = &cli.Command{
		Name:   "make:handler",
		Usage:  "Creates a handler file",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The handler name, (i.e: "healthcheck")`,
			},
		},
	}
	return &cc
}

func (cc *MakeHandlerCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()
	name := c.String("name")

	if len(strings.Trim(shortcut, " ")) > 0 {
		name = shortcut
	} else {
		if len(name) == 0 {
			fmt.Println("\nPlease provide the handler name, for example: --name healthcheck")
			return nil
		}
	}

	return cc.Generate(name)
}

func (cc *MakeHandlerCommand) Generate(name string) error {
	files := map[string]string{
		path.Load().HandlersPath(strcase.ToSnake(name) + ".go"): "stubs/handler/single.stub",
	}

	fmt.Println("Created handler, located at:")

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)
		content = strings.Replace(content, "##CAMEL_CASE_NAME##", strcase.ToCamel(name), -1)
		content = strings.Replace(content, "##SNAKE_CASE_NAME##", strcase.ToSnake(name), -1)

		//> create a file and write the content
		err = php.FilePutContents(orig, content, 0755)

		if err != nil {
			return err
		}

		fmt.Printf(" > %s\n", orig)
	}

	fmt.Println("\nGo to registrar/routes.go and paste this:")
	fmt.Println()
	fmt.Println("    var Routes = &[]routes.Routing{")
	fmt.Println("    	...,")
	fmt.Printf(`        {
		Path:    "/%s",
		Name:    "%s",
		Method:  routes.Method{"GET"}, // defaulting to "GET"
		Handler: handlers.%s,
	},`,
		strcase.ToKebab(name),
		strcase.ToKebab(name),
		strcase.ToCamel(name),
	)
	fmt.Println("\n    }")

	return nil
}
