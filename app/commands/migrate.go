package commands

import (
	"fmt"

	"github.com/lucidfy/lucid/databases"
	"github.com/lucidfy/lucid/databases/migrations"
	cli "github.com/urfave/cli/v2"
)

type MigrateCommand struct {
	Command *cli.Command
}

func Migrate() *MigrateCommand {
	var cc MigrateCommand
	cc.Command = &cli.Command{
		Name:   "migrate",
		Usage:  "Calls gorm's AutoMigrate",
		Action: cc.Handle,
	}
	return &cc
}

func (cc *MigrateCommand) Handle(c *cli.Context) error {
	db := databases.Resolve()

	for _, model := range migrations.AutoMigrate {
		db.AutoMigrate(model)

		fmt.Println("")

		i, ok := model.(interface{ TableName() string })
		if ok {
			fmt.Printf("Migrated %q table\n", i.TableName())
		}
	}

	return nil
}
