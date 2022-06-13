package users

import (
	"database/sql"
	"time"

	"github.com/golang-module/carbon"
)

const Table = "users"

var Updatables = []string{
	"name",
	"email",
	"password",
}

type Model struct {
	ID              uint           `gorm:"primarykey;auto_increment;not_null" json:"id"`
	Name            string         `gorm:"column:name" json:"name"`
	Email           string         `gorm:"column:email" json:"email"`
	EmailVerifiedAt *time.Time     `gorm:"column:email_verified_at" json:"-"`
	Password        string         `gorm:"column:password" json:"-"`
	RememberToken   sql.NullString `gorm:"column:remember_token" json:"-"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func (Model) TableName() string { // https://gorm.io/docs/conventions.html#TableName
	return Table
}

func (m Model) ReadableCreatedAt() string {
	return carbon.Time2Carbon(m.CreatedAt).ToDayDateTimeString()
}

func (m Model) ReadableUpdatedAt() string {
	return carbon.Time2Carbon(m.UpdatedAt).ToDayDateTimeString()
}
