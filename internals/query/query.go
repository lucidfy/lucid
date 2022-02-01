package query

import (
	"database/sql"
	"reflect"
)

func Select(db *sql.DB, model interface{}, stmt string, args ...interface{}) []interface{} {
	rows, err := db.Query(stmt, args...)

	if err != nil {
		panic(err)
	}

	// Get the model attributes
	// -> by getting the number of fields as the number of columns
	// -> then we will make an interface out of the columns
	s := reflect.ValueOf(model).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)

	var records []interface{}

	defer rows.Close()

	for rows.Next() {
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)

		if err != nil {
			panic(err)
		}

		records = append(records, model)
	}

	return records
}
