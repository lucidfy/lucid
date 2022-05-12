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

type CraftValidationCommand struct {
	Command *cli.Command
}

func CraftValidation() *CraftValidationCommand {
	var cc CraftValidationCommand
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

func (cc *CraftValidationCommand) Handle(c *cli.Context) error {
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

func (cc *CraftValidationCommand) Generate(name string) error {
	files := map[string]string{
		path.Load().BasePath("app/validations/" + strcase.ToSnake(name) + ".go"): "stubs/validation.stub",
	}

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
		err = php.FilePutContents(orig, content, 0755)

		if err != nil {
			return err
		}
	}

	fmt.Println("\nSuccessfully created a validation!")

	return nil
}
