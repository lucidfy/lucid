package users

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/daison12006013/gorvel/databases"
)

func Lists(s *SearchableTable) error {
	db := databases.Resolve()

	var total int
	countBuilder := sq.Select("count(*)").From(Table)
	countStmt, countArgs, _ := query(s, countBuilder).ToSql()
	db.Raw(countStmt, countArgs...).Scan(&total)

	var records []Attributes
	selectBuilder := sq.Select("*").From(Table).
		OrderBy(*s.OrderByCol + " " + *s.OrderBySort).
		Limit(uint64(s.Paginate.PerPage)).
		Offset(uint64(((s.Paginate.CurrentPage) - 1) * s.Paginate.PerPage))

	selectStmt, selectArgs, _ := query(s, selectBuilder).ToSql()
	db.Raw(selectStmt, selectArgs...).Scan(&records)

	// reconstruct the searchable table
	s.Paginate.Reconstruct(&records, total)

	return nil
}

func query(s *SearchableTable, builder sq.SelectBuilder) sq.SelectBuilder {
	for _, header := range s.Headers {
		if !header.Input.CanSearch || header.Input.Value == "" {
			continue
		}

		var pred sq.Or
		for _, searchColumn := range header.Input.SearchColumn {
			switch header.Input.SearchPattern {
			case "-":
				pred = append(pred, sq.Eq{searchColumn: fmt.Sprintf("%v", header.Input.Value)})
			case "<-":
				pred = append(pred, sq.Like{searchColumn: "%" + fmt.Sprintf("%v", header.Input.Value)})
			case "->":
				pred = append(pred, sq.Like{searchColumn: fmt.Sprintf("%v", header.Input.Value) + "%"})
			case "<->":
				pred = append(pred, sq.Like{searchColumn: "%" + fmt.Sprintf("%v", header.Input.Value) + "%"})
			}
		}
		builder = builder.Where(pred)
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
