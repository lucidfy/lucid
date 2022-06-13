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

type MakeMigratorCommand struct {
	Command *cli.Command
}

func MakeMigrator() *MakeMigratorCommand {
	var cc MakeMigratorCommand
	cc.Command = &cli.Command{
		Name:   "make:migrator",
		Usage:  "Creates a migrator",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "model",
				Value: "",
				Usage: `The migrator model (i.e: "reports")`,
			},
		},
	}
	return &cc
}

func (cc *MakeMigratorCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()

	model := c.String("model")

	if len(strings.Trim(shortcut, " ")) > 0 {
		model = shortcut
	} else {
		if len(model) == 0 {
			fmt.Println("\nPlease provide the filename, for example: --model reports")
			return nil
		}
	}

	return cc.Generate(model)
}

func (cc *MakeMigratorCommand) Generate(model string) error {
	files := map[string]string{
		path.Load().DatabasePath("migrations/" + model + "_migrator.go"): "stubs/make_migrator.stub",
	}

	fmt.Println("Created migrator, located at:")

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)
		content = strings.Replace(content, "##SNAKE_CASE_NAME##", strcase.ToSnake(model), -1)
		content = strings.Replace(content, "##CAMEL_CASE_NAME##", strcase.ToCamel(model), -1)

		//> create a file and write the content
		app_err := php.FilePutContents(orig, content, 0755)
		if app_err != nil {
			return app_err.Error
		}

		fmt.Printf(" > %s\n", orig)
	}

	fmt.Println("\nGo to registrar/migrations.go and copy below line:")
	fmt.Println()
	fmt.Println("    var Migrations = []interface{}{")
	fmt.Println("    	...,")
	fmt.Println("       &migrations." + strcase.ToCamel(model) + "Migration{},")
	fmt.Println("    }")

	return nil
}
