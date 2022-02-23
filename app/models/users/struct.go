package users

import (
	"database/sql"
	"time"
)

const Table = "users"
const PrimaryKey = "id"

type Attributes struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	Name            string         `gorm:"column:name" json:"name"`
	Email           string         `gorm:"column:email" json:"email"`
	EmailVerifiedAt *time.Time     `gorm:"column:email_verified_at" json:"email_verified_at,omitempty"`
	Password        string         `gorm:"column:password" json:"password"`
	RememberToken   sql.NullString `gorm:"column:remember_token" json:"remember_token,omitempty"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
}
