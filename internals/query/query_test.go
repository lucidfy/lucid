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
	var user UserModel

	db, err := sql.Open("sqlite3", "../../databases/sqlite.db")

	if err != nil {
		panic(err)
	}

	schema := Interpreter()
	stmt := schema.Table("users").Where("id = ?", 1).Limit(1).ToSql()

	records := Select(
		db,
		&user,
		stmt,
		schema.GetBindings()...,
	)

	for _, u := range records {
		record := u.(*UserModel) // convert interface into the Struct

		got := record.Name
		want := "John Doe"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}
