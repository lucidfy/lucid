package users

import (
	"database/sql"
	"os"
	"time"

	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/path"
	"github.com/daison12006013/gorvel/pkg/paginate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const Table = "users"
const PrimaryKey = "id"

type Paginate struct {
	paginate.Paginate
	Items []Attributes
}

type Attributes struct {
	gorm.Model
	ID              uint           `json:"id" gorm:"primarykey"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	Password        string         `json:"password"`
	RememberToken   sql.NullString `json:"remember_token"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func db() *gorm.DB {
	filepath := path.Load().DatabasePath(os.Getenv("DB_DATABASE"))
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})

	if err != nil {
		logger.Fatal(err)
	}

	return db
}
