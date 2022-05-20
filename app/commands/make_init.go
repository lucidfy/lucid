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

type MakeInitCommand struct {
	Command *cli.Command
}

func MakeInit() *MakeInitCommand {
	var cc MakeInitCommand
	cc.Command = &cli.Command{
		Name:   "make:init",
		Usage:  "Initializes a new crafter",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The crafter name, (i.e: "middleware")`,
			},
		},
	}
	return &cc
}

func (cc *MakeInitCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()
	name := c.String("name")

	if len(strings.Trim(shortcut, " ")) > 0 {
		name = shortcut
	} else {
		if len(name) == 0 {
			fmt.Println("\nPlease provide the crafter name, for example: --name middleware")
			return nil
		}
	}

	return cc.Generate(name)
}

func (cc *MakeInitCommand) Generate(name string) error {
	files := map[string]string{
		path.Load().ConsolePath(strcase.ToSnake("make_"+name) + ".go"): "stubs/make_init.stub",
	}

	fmt.Println("Created a crafter, located at:")

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

		// replace those string replacers back to their original key
		content = strings.Replace(content, "--SNAKE_CASE_NAME--", "##SNAKE_CASE_NAME##", -1)
		content = strings.Replace(content, "--CAMEL_CASE_NAME--", "##CAMEL_CASE_NAME##", -1)

		//> create a file and write the content
		app_err := php.FilePutContents(orig, content, 0755)
		if app_err != nil {
			return app_err.Error
		}

		fmt.Printf(" > %s\n", orig)
	}

	fmt.Println("")

	return nil
}
