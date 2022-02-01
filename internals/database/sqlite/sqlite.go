package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/daison12006013/gorvel/internals/query"
	_ "github.com/mattn/go-sqlite3"
)

type Driver struct {
	Database *sql.DB
}

func Make(dbname string) Driver {
	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		panic(err)
	}

	driver := Driver{db}

	return driver
}

func (d Driver) Query(stmt string, model interface{}) []interface{} {
	records := query.Select(d.Database, model, stmt)

	if len(records) > 0 {
		return records
	}

	return nil
}

func (d Driver) First(table string, primaryKey string, model interface{}) interface{} {
	stmt := fmt.Sprintf("SELECT * FROM %s order by %s asc limit 1", table, primaryKey)

	records := query.Select(d.Database, model, stmt)

	if len(records) > 0 {
		return records[0]
	}

	return nil
}

func (d Driver) Last(table string, primaryKey string, model interface{}) interface{} {
	stmt := fmt.Sprintf("SELECT * FROM %s order by %s desc limit 1", table, primaryKey)

	records := query.Select(d.Database, model, stmt)

	if len(records) > 0 {
		return records[0]
	}

	return nil
}
