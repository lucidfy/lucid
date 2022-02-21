package users

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/daison12006013/gorvel/databases"
)

func Lists(p *Paginate) error {
	db := databases.Resolve()

	countBuilder := sq.Select("count(*)").From(Table)
	selectBuilder := sq.Select("*").From(Table).
		OrderBy(*p.OrderByCol + " " + *p.OrderBySort).
		Limit(uint64(p.PerPage)).
		Offset(uint64(((p.CurrentPage) - 1) * p.PerPage))

	countStmt, countArgs, _ := query(p, countBuilder).ToSql()
	selectStmt, selectArgs, _ := query(p, selectBuilder).ToSql()

	var total int
	db.Raw(countStmt, countArgs...).Scan(&total)

	var records []Attributes
	db.Raw(selectStmt, selectArgs...).Scan(&records)

	p.Reconstruct(&records, total)
	return nil
}

func query(p *Paginate, builder sq.SelectBuilder) sq.SelectBuilder {
	if p.TextSearch != nil {
		builder = builder.Where(sq.Like{"name": *p.TextSearch + "%"})
	}
	return builder
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
