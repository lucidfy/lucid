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

type MakeValidationCommand struct {
	Command *cli.Command
}

func MakeValidation() *MakeValidationCommand {
	var cc MakeValidationCommand
	cc.Command = &cli.Command{
		Name:   "make:validation",
		Usage:  "Creates a validation file",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The validation name, (i.e: "reports")`,
			},
		},
	}
	return &cc
}

func (cc *MakeValidationCommand) Handle(c *cli.Context) error {
	shortcut := c.Args().First()
	name := c.String("name")

	if len(strings.Trim(shortcut, " ")) > 0 {
		name = shortcut
	} else {
		if len(name) == 0 {
			fmt.Println("\nPlease provide the validation name, for example: --name reports")
			return nil
		}
	}

	return cc.Generate(name)
}

func (cc *MakeValidationCommand) Generate(name string) error {
	files := map[string]string{
		path.Load().BasePath("app/validations/" + strcase.ToSnake(name) + ".go"): "stubs/validation.stub",
	}

	fmt.Println("Created validation, located at:")

	for orig, stub := range files {
		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)
		content = strings.Replace(content, "##CAMEL_CASE_NAME##", strcase.ToCamel(name), -1)

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
