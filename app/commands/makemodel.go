package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	cli "github.com/urfave/cli/v2"
)

type CraftModelCommand struct {
	Command *cli.Command
}

func CraftModel() *CraftModelCommand {
	var cc CraftModelCommand
	cc.Command = &cli.Command{
		Name:    "make:model",
		Aliases: []string{"m:m"},
		Usage:   "Creates a model",
		Action:  cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The package name (i.e: "profiles")`,
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

func (cc *CraftModelCommand) Handle(c *cli.Context) error {
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

func (cc *CraftModelCommand) Generate(name string, table string) error {
	files := map[string]string{
		"model_test.go": "stubs/models/model_test.stub",
		"model.go":      "stubs/models/model.stub",
		"struct.go":     "stubs/models/struct.stub",
	}

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := strings.Replace(string(stubContent), "##TABLE_NAME##", table, -1)
		content = strings.Replace(content, "##PACKAGE_NAME##", table, -1)

		//> create the directory
		err = php.Mkdir(
			path.Load().ModelsPath(name),
			os.ModePerm,
			true,
		)

		if err != nil {
			return err
		}

		//> create a file and write the content
		err = php.FilePutContents(
			path.Load().ModelsPath(name+"/"+orig),
			content,
			0755,
		)

		if err != nil {
			return err
		}
	}

	fmt.Println("\nSuccessfully created a model!")

	return nil
}
