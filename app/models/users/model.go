package users

import (
	"math"
	"os"

	"github.com/daison12006013/gorvel/pkg/dbadapter"
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
	var err error

	db := db()
	selectStmt := query.Interpreter().
		Table(Table).
		OrderBy(orderByCol, orderBySort).
		Limit(perPage).
		Offset((currentPage - 1) * perPage).
		ToSql()

	countStmt := query.Interpreter().
		Table(Table).
		CountSql()

	// query the total count
	var total int
	err = db.Select(countStmt).Find(&total)
	if err != nil {
		return nil, err
	}

	// query the records
	var records []Attributes
	err = db.Select(selectStmt).Find(&records)
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
	db := db()
	interpreter := query.Interpreter()
	selectStmt := interpreter.
		Table(Table).
		Where("id = ?", id).
		Limit(1).
		ToSql()

	var records []Attributes

	err := db.Select(selectStmt).Find(
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
