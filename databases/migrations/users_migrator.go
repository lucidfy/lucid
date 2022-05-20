package migrations

import (
	"github.com/lucidfy/lucid/app/models/users"
	"gorm.io/gorm"
)

type UserMigration struct{}

// Auto, to migrate the table / columns automatically
// by running `./run migrate --auto-migrate`
func (u UserMigration) Auto(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&users.Model{})

	return err
}

// Up, to increment your migration
func (u UserMigration) Up(m gorm.Migrator) (err error) {
	if !u.model_exists(m) {
		err = m.CreateTable(&users.Model{})
	}

	return err
}

// Down, to drop a group of migrations
func (u UserMigration) Down(m gorm.Migrator) (err error) {
	if u.model_exists(m) {
		err = m.DropTable(&users.Model{})
	}

	return err
}

func (u UserMigration) model_exists(m gorm.Migrator) bool {
	return m.HasTable(&users.Model{})
}
