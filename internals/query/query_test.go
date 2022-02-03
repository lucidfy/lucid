package query

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type UserModel struct {
	ID              int     `json:"id" db:"id"`
	Name            string  `json:"name" db:"name"`
	Email           string  `json:"email" db:"email"`
	EmailVerifiedAt string  `json:"email_verified_at" db:"email_verified_at"`
	Password        string  `json:"password" db:"password"`
	RememberToken   *string `json:"remember_token" db:"remember_token"`
	CreatedAt       string  `json:"created_at" db:"created_at"`
	UpdatedAt       string  `json:"updated_at" db:"updated_at"`
}

func TestQuerySelect(t *testing.T) {
	var users []UserModel

	driver, err := sql.Open("sqlite3", "../../databases/sqlite.db")

	if err != nil {
		t.Errorf("Could not connect to sqlite: %q", err)
	}

	interpreter := Interpreter()
	stmt := interpreter.Table("users").Where("id = ?", 1).Limit(1).ToSql()

	err = Connect(driver).
		Select(stmt).
		Find(&users, interpreter.GetBindings()...)

	if err != nil {
		t.Errorf("query throws an error %q", err)
	}

	for _, record := range users {
		got := record.Name
		want := "John Doe"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
