package users

import (
	"github.com/Masterminds/squirrel"
	"github.com/daison12006013/gorvel/databases"
)

func Lists(baseUrl string, currentPage int, perPage int, orderByCol string, orderBySort string) (*Paginate, error) {
	db := databases.Resolve()

	selectStmt, _, _ := squirrel.
		Select("*").
		From(Table).
		OrderBy(orderByCol + " " + orderBySort).
		Limit(uint64(perPage)).
		Offset(uint64(((currentPage) - 1) * perPage)).
		ToSql()

	countStmt, _, _ := squirrel.Select("count(*)").From(Table).ToSql()

	var total int
	db.Raw(countStmt).Scan(&total)

	var records []Attributes
	db.Raw(selectStmt).Scan(&records)

	var paginated Paginate
	paginated.BaseUrl = baseUrl
	paginated.Reconstruct(&records, total, perPage, currentPage)

	return &paginated, nil
}

// func FindById(id string) (*Attributes, error) {
// 	conn := db()
// 	interpreter := query.Interpreter()
// 	selectStmt := interpreter.
// 		Table(Table).
// 		Where("id = ?", id).
// 		Limit(1).
// 		ToSql()

// 	var records []Attributes

// 	err := conn.Select(selectStmt).Find(
// 		&records,
// 		interpreter.GetBindings()...,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(records) > 0 {
// 		return &records[0], nil
// 	}

// 	return nil, nil
// }

// func Exists(id string) bool {
// 	conn := db()
// 	interpreter := query.Interpreter()
// 	countStmt := interpreter.
// 		Table(Table).
// 		Where("id = ?", id).
// 		CountSql()

// 	var total int
// 	err := conn.Select(countStmt).Find(
// 		&total,
// 		interpreter.GetBindings()...,
// 	)
// 	if err != nil {
// 		logger.Fatal(err)
// 		return false
// 	}

// 	return total > 0
// }

// func DeleteById(id string) bool {
// 	conn := db()
// 	stmt := fmt.Sprintf("DELETE FROM %s where id = $1;", Table)
// 	_, err := conn.DB.Exec(stmt, id)
// 	if err != nil {
// 		logger.Fatal(err)
// 		return false
// 	}

// 	return true
// }
