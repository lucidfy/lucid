package commands

import (
	"fmt"

	"github.com/lucidfy/lucid/databases"
	cli "github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func Migrate(c *cli.Context, migrations []interface{}) error {
	db := databases.Resolve()

	if c.Bool("current-database") {
		fmt.Println(db.Migrator().CurrentDatabase())
		return nil
	}

	if c.Bool("auto-migrate") {
		for _, migration := range migrations {
			i, ok := migration.(interface {
				Auto(*gorm.DB) error
			})

			if ok {
				if err := i.Auto(db); err == nil {
					fmt.Printf("Auto Migrated %T\n", migration)
				}
			}
		}
		return nil
	}

	if c.Bool("up") {
		for _, migration := range migrations {
			i, ok := migration.(interface {
				Up(gorm.Migrator) error
			})

			if ok {
				if err := i.Up(db.Migrator()); err == nil {
					fmt.Printf("Migrated %T\n", migration)
				}
			}
		}
	}

	if c.Bool("down") {
		for _, migration := range migrations {
			i, ok := migration.(interface {
				Down(gorm.Migrator) error
			})

			if ok {
				if err := i.Down(db.Migrator()); err == nil {
					fmt.Printf("Rollback %T\n", migration)
				}
			}
		}
	}

	return nil
}
