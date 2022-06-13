package lucid_commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	cli "github.com/urfave/cli/v2"
)

type MakeModelCommand struct {
	Command *cli.Command
}

func MakeModel() *MakeModelCommand {
	var cc MakeModelCommand
	cc.Command = &cli.Command{
		Name:   "make:model",
		Usage:  "Creates a model",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The package name (i.e: "userfriends")`,
			},
			&cli.StringFlag{
				Name:  "table",
				Value: "",
				Usage: `The name of your table (i.e: "user_friends")`,
			},
		},
	}
	return &cc
}

func (cc *MakeModelCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()

	name := c.String("name")
	table := c.String("table")

	if len(strings.Trim(shortcut, " ")) > 0 {
		name = shortcut
		table = shortcut
	} else {
		if len(name) == 0 {
			fmt.Println("\nPlease provide the filename, for example: --name profiles")
			return nil
		}

		if len(table) == 0 {
			fmt.Println("\nPlease provide the table to use, for example: --table profiles")
			return nil
		}
	}

	return cc.Generate(name, table)
}

func (cc *MakeModelCommand) Generate(name string, table string) error {
	folder := strcase.ToSnake(name)
	files := map[string]string{
		path.Load().ModelsPath(folder + "/model_test.go"): "stubs/models/model_test.stub",
		path.Load().ModelsPath(folder + "/model.go"):      "stubs/models/model.stub",
		path.Load().ModelsPath(folder + "/struct.go"):     "stubs/models/struct.stub",
	}

	//> create the directory
	app_err := php.Mkdir(
		path.Load().ModelsPath(folder),
		os.ModePerm,
		true,
	)

	if app_err != nil {
		return app_err.Error
	}

	fmt.Println("Created model, located at:")

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)
		content = strings.Replace(content, "##TABLE_NAME##", table, -1)
		content = strings.Replace(content, "##PACKAGE_NAME##", strcase.ToSnake(table), -1)

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
