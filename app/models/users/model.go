package users

import (
	"github.com/daison12006013/gorvel/databases"
	"github.com/daison12006013/gorvel/pkg/paginate/searchable"
)

func Lists(st *searchable.Table) error {
	db := databases.Resolve()

	// fetch counts
	var total int
	countStmt, countArgs, _ := st.QueryCount(Table).ToSql()
	db.Raw(countStmt, countArgs...).Scan(&total)

	// fetch the data
	var records []Attributes
	selectStmt, selectArgs, _ := st.QuerySelect(Table).ToSql()
	db.Raw(selectStmt, selectArgs...).Scan(&records)

	// reload the pagination data
	st.Paginate.Reconstruct(&records, total)

	return nil
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
