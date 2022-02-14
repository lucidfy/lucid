package users

import (
	"fmt"
	"math"
	"os"

	"github.com/daison12006013/gorvel/pkg/dbadapter"
	"github.com/daison12006013/gorvel/pkg/facade/logger"
	"github.com/daison12006013/gorvel/pkg/facade/path"
	"github.com/daison12006013/gorvel/pkg/query"
)

// connect to our database
func db() *query.Result {
	db := query.Connect(dbadapter.SQLite(
		path.Load().DatabasePath(os.Getenv("DB_DATABASE")),
	))
	return db
}

func Lists(currentPage int, perPage int, orderByCol string, orderBySort string) (*Paginate, error) {
	conn := db()

	selectStmt := query.Interpreter().
		Table(Table).
		OrderBy(orderByCol, orderBySort).
		Limit(perPage).
		Offset((currentPage - 1) * perPage).
		ToSql()

	countStmt := query.Interpreter().
		Table(Table).
		CountSql()

	var total int

	err := conn.Select(countStmt).Find(&total)
	if err != nil {
		return nil, err
	}

	// query the records
	var records []Attributes
	err = conn.Select(selectStmt).Find(&records)
	if err != nil {
		return nil, err
	}

	var paginated Paginate
	paginated.PerPage = perPage
	paginated.CurrentPage = currentPage
	paginated.Data = records
	paginated.Total = total
	paginated.LastPage = int(math.Ceil(float64(total) / float64(perPage)))

	return &paginated, nil
}

func FindById(id string) (*Attributes, error) {
	conn := db()
	interpreter := query.Interpreter()
	selectStmt := interpreter.
		Table(Table).
		Where("id = ?", id).
		Limit(1).
		ToSql()

	var records []Attributes

	err := conn.Select(selectStmt).Find(
		&records,
		interpreter.GetBindings()...,
	)

	if err != nil {
		return nil, err
	}

	if len(records) > 0 {
		return &records[0], nil
	}

	return nil, nil
}

func Exists(id string) bool {
	conn := db()
	interpreter := query.Interpreter()
	countStmt := interpreter.
		Table(Table).
		Where("id = ?", id).
		CountSql()

	var total int
	err := conn.Select(countStmt).Find(
		&total,
		interpreter.GetBindings()...,
	)
	if err != nil {
		logger.Fatal(err)
		return false
	}

	return total > 0
}

func DeleteById(id string) bool {
	conn := db()
	stmt := fmt.Sprintf("DELETE FROM %s where id = $1;", Table)
	_, err := conn.DB.Exec(stmt, id)
	if err != nil {
		logger.Fatal(err)
		return false
	}

	return true
}
