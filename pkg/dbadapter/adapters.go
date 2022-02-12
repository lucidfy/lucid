package dbadapter

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func MySQL() {
	// * TODO
}

func PostgreSQL() {
	// * TODO
}

func SQLite(dbname string) *sql.DB {
	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		panic(err)
	}

	return db
}
