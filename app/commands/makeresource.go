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

type CraftResourceCommand struct {
	Command *cli.Command
}

func CraftResource() *CraftResourceCommand {
	var cc CraftResourceCommand
	cc.Command = &cli.Command{
		Name:   "make:resource",
		Usage:  "Creates a resource along with the model",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The package name, (i.e: "reports")`,
			},
			&cli.StringFlag{
				Name:  "table",
				Value: "",
				Usage: `The name of your table (i.e: "user_profiles")`,
			},
		},
	}
	return &cc
}

func (cc *CraftResourceCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()
	name := c.String("name")
	table := c.String("table")

	if len(strings.Trim(shortcut, " ")) > 0 {
		name = shortcut
		table = shortcut
	} else {
		if len(name) == 0 {
			fmt.Println("\nPlease provide the resource name, for example: --name reports")
			return nil
		}

		if len(table) == 0 {
			fmt.Println("\nPlease provide the table to use, for example: --table profiles")
			return nil
		}
	}

	packageName := strcase.ToSnake(name + "_handler")
	smallCaseName := strcase.ToSnake(name)
	camelCaseName := strcase.ToCamel(name)

	resp := cc.Generate(packageName, smallCaseName, camelCaseName)

	model := CraftModelCommand{}
	model.Generate(name, table)

	validation := CraftValidationCommand{}
	validation.Generate(name)

	return resp
}

func (cc *CraftResourceCommand) Generate(packageName string, smallCaseName string, camelCaseName string) error {
	files := map[string]string{
		path.Load().HandlersPath(packageName + "/create.go"): "stubs/handler/resource/create.stub",
		path.Load().HandlersPath(packageName + "/delete.go"): "stubs/handler/resource/delete.stub",
		path.Load().HandlersPath(packageName + "/lists.go"):  "stubs/handler/resource/lists.stub",
		path.Load().HandlersPath(packageName + "/update.go"): "stubs/handler/resource/update.stub",
	}

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)
		content = strings.Replace(content, "##PACKAGE_NAME##", packageName, -1)
		content = strings.Replace(content, "##SMALL_CASE_NAME##", smallCaseName, -1)
		content = strings.Replace(content, "##CAMEL_CASE_NAME##", camelCaseName, -1)

		//> create the directory
		err = php.Mkdir(
			path.Load().HandlersPath(packageName),
			os.ModePerm,
			true,
		)

		if err != nil {
			return err
		}

		//> create a file and write the content
		err = php.FilePutContents(orig, content, 0755)

		if err != nil {
			return err
		}
	}

	fmt.Println("\nSuccessfully created a resource handlers!")

	return nil
}
