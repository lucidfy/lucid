package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func CraftHandler(c *cli.Context) error {
	fmt.Println("make handler ", c.Args().First(), c.String("name"))
	return nil
}
