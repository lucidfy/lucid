package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucidfy/lucid/pkg/facade/logger"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"github.com/lucidfy/lucid/pkg/functions/php"
	cli "github.com/urfave/cli/v2"
)

func CraftModel(c *cli.Context) error {
	shortcut := c.Args().First()
	name := c.String("name")
	table := c.String("table")
	if len(shortcut) != 0 {
		name = shortcut
		table = shortcut
	}

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
			panic(err)
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
			logger.Error("Mkdir Error ", err)
			panic(err)
		}

		//> create a file and write the content
		err = php.FilePutContents(
			path.Load().ModelsPath(name+"/"+orig),
			content,
			0775,
		)

		if err != nil {
			logger.Error("FilePutContents Error ", err)
			panic(err)
		}
	}

	fmt.Println("\nSuccessfully created a model!")

	return nil
}
