package lucid_commands

import (
	"fmt"
	"os"

	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	cli "github.com/urfave/cli/v2"
)

type MakeTinkerCommand struct {
	Command *cli.Command
}

func MakeTinker() *MakeTinkerCommand {
	var cc MakeTinkerCommand
	cc.Command = &cli.Command{
		Name:   "tinker",
		Usage:  "Creates a tinker_test.go in the root",
		Action: cc.Handle,
	}
	return &cc
}

func (cc *MakeTinkerCommand) Handle(c *cli.Context) error {
	return cc.Generate()
}

func (cc *MakeTinkerCommand) Generate() error {
	files := map[string]string{
		path.Load().BasePath("tinker_test.go"): "stubs/tinker_test.stub",
	}

	for orig, stub := range files {
		if php.FileExists(orig) {
			fmt.Printf("File [%s] already exists!", orig)
			continue
		}

		//> read the stub and parse it
		//> then replace all the keys
		stubContent, err := os.ReadFile(stub)
		if err != nil {
			return err
		}

		content := string(stubContent)

		//> create a file and write the content
		app_err := php.FilePutContents(orig, content, 0755)
		if app_err != nil {
			return app_err.Error
		}

		fmt.Printf(" > %s\n", orig)
	}

	fmt.Println("\n\nStart tinkering!")

	return nil
}
