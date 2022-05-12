package commands

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

type CraftHandlerCommand struct {
	Command *cli.Command
}

func CraftHandler() *CraftHandlerCommand {
	var cc CraftHandlerCommand
	cc.Command = &cli.Command{
		Name:   "make:handler",
		Usage:  "Creates an http handler",
		Action: cc.Handle,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: `The handler's name, (i.e: "users")`,
			},
		},
	}
	return &cc
}

func (cc *CraftHandlerCommand) Handle(c *cli.Context) error {
	fmt.Println("make handler ", c.Args().First(), c.String("name"))
	return nil
}
