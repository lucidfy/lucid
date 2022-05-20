package migrations_model

import (
	"database/sql"
	"time"

	"github.com/golang-module/carbon"
)

const Table = "migrations"
const PrimaryKey = "id"

var Updatables = []string{
	"migrator",
	"sequence",
}

type Model struct {
	ID        uint          `gorm:"primarykey;auto_increment;not_null" json:"id"`
	Name      string        `gorm:"column:name"`
	Sequence  sql.NullInt16 `gorm:"column:sequence"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (Model) TableName() string {
	return Table
}

func (m Model) ReadableCreatedAt() string {
	return carbon.Time2Carbon(m.CreatedAt).ToDayDateTimeString()
}

func (m Model) ReadableUpdatedAt() string {
	return carbon.Time2Carbon(m.UpdatedAt).ToDayDateTimeString()
}
