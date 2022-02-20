package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func CraftResource(c *cli.Context) error {
	fmt.Println("make resources ", c.Args().First())
	return nil
}
