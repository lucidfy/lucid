package migrations_model

import (
	"database/sql"
	"time"
)

const Table = "migrations"

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
