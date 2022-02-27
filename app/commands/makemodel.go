package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func CraftModel(c *cli.Context) error {
	fmt.Println("make model ", c.Args().First(), c.String("name"), c.String("table"))
	return nil
}
