package commands

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lucidfy/lucid/databases"
	"github.com/lucidfy/lucid/pkg/models/migrations_model"
	cli "github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type MigrateCommand struct {
	Command    *cli.Context
	Migrations []interface{}
}

func (mc MigrateCommand) Handle() error {
	c := mc.Command
	migrations := mc.Migrations

	db := databases.Resolve()
	db.AutoMigrate(&migrations_model.Model{})

	if c.Bool("current-database") {
		fmt.Println(db.Migrator().CurrentDatabase())
		return nil
	}

	if c.Bool("auto") {
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
				struct_name := fmt.Sprintf("%T", migration)
				if !mc.CheckStructNameIfExists(struct_name) {
					if err := i.Up(db.Migrator()); err == nil {
						fmt.Printf("Migrated %T\n", migration)
					}

					mc.Save(struct_name, mc.GetCurrentSequence()+1)
				}
			}
		}
		return nil
	}

	if c.Bool("down") {
		for _, migration := range migrations {
			i, ok := migration.(interface {
				Down(gorm.Migrator) error
			})

			if ok {
				struct_name := fmt.Sprintf("%T", migration)
				if mc.CheckStructNameIfExists(struct_name) {
					if !mc.CheckStructNameOnCurrentSequence(struct_name) {
						continue
					}

					if err := i.Down(db.Migrator()); err == nil {
						fmt.Printf("Rollback %T\n", migration)
					}

					mc.Delete(struct_name)
				}
			}
		}
		return nil
	}

	return nil
}

// CheckStructNameIfExists, check if the struct name ever existed
func (mc MigrateCommand) CheckStructNameIfExists(struct_name string) bool {
	db := databases.Resolve()

	stmt, args, _ := sq.Select("1").From(migrations_model.Table).Where(sq.Eq{"name": struct_name}).ToSql()
	stmt, args, _ = sq.Expr("select exists("+stmt+") as found", args).ToSql()
	var found bool
	db.Raw(stmt, args).Scan(&found)

	return found
}

// CheckStructNameOnCurrentSequence
// we check if the struct name is on the current sequence
func (mc MigrateCommand) CheckStructNameOnCurrentSequence(struct_name string) bool {
	db := databases.Resolve()
	stmt, args, _ := sq.Select("1").
		From(migrations_model.Table).
		Where(sq.Eq{"name": struct_name}).
		Where(sq.Eq{"sequence": mc.GetCurrentSequence()}).
		ToSql()
	var found bool
	db.Raw(stmt, args...).Scan(&found)

	return found
}

// Save, stores the struct name inside migrations table
func (mc MigrateCommand) Save(struct_name string, seq uint) {
	db := databases.Resolve()

	db.Create(&migrations_model.Model{
		Name: struct_name,
		Sequence: sql.NullInt16{
			Int16: int16(seq),
			Valid: true,
		},
	})
}

// Delete, deletes the struct name inside migrations table
func (mc MigrateCommand) Delete(struct_name string) {
	db := databases.Resolve()
	db.Where("name = ?", struct_name).Delete(&migrations_model.Model{})
}

// GetCurrentSequence, gets the current sequence in the database
// if there is no record inside migrations table, the default will be 0
func (mc MigrateCommand) GetCurrentSequence() uint {
	db := databases.Resolve()

	stmt := fmt.Sprintf(
		`select sequence from %s order by id desc limit 1`,
		migrations_model.Table,
	)

	var seq sql.NullInt16
	db.Raw(stmt).Scan(&seq)

	val, err := seq.Value()
	if err != nil {
		panic(err)
	}

	if !seq.Valid || val == nil {
		return 0
	}

	return uint(val.(int64))
}
