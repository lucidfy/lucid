package lucid_commands

import (
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	MakeInit().Command,
	MakeHandler().Command,
	MakeResource().Command,
	MakeModel().Command,
	MakeValidation().Command,
}
