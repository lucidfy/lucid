package query

import (
	"database/sql"
	"errors"
	"reflect"
)

type Result struct {
	db        *sql.DB
	statement string
	data      []interface{}
	err       error
}

func Connect(db *sql.DB) *Result {
	r := Result{db, "", nil, nil}

	return &r
}

func (r *Result) Get() ([]interface{}, error) {
	return r.data, r.err
}

func (r *Result) Select(stmt string) *Result {
	r.statement = stmt

	return r
}

func (r *Result) Find(model interface{}, args ...interface{}) error {
	rows, err := r.db.Query(r.statement, args...)

	if err != nil {
		return err
	}

	defer rows.Close()

	// modelv is the model slice
	s := reflect.ValueOf(model)

	var numCols int
	var columns []interface{}
	var values reflect.Value
	var val reflect.Value
	var rowt reflect.Type

	switch reflect.TypeOf(model).Elem().Kind() {
	case reflect.Slice:
		values = s.Elem()
		rowt = values.Type().Elem()
		val = reflect.New(rowt).Elem()
		numCols = rowt.NumField()
		columns = make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			columns[i] = val.
				FieldByName(rowt.Field(i).Name).
				Addr().
				Interface()
		}
	default:
		values = s.Elem()
		rowt = values.Type()
		val = reflect.New(rowt).Elem()
		selectCols, _ := rows.Columns()
		numCols = len(selectCols)
		columns = make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			columns[i] = val.Addr().Interface()
		}
	}

	for rows.Next() {
		err = rows.Scan(columns...)

		if err != nil {
			panic(err)
		}

		// Append struct to result slice. Because the struct
		// is copied in append, we can reuse the struct in
		// this loop.
		switch reflect.TypeOf(model).Elem().Kind() {
		case reflect.Slice:
			values.Set(reflect.Append(values, val))
		default:
			values.Set(val)
		}
	}

	return nil
}

func (r *Result) First() (interface{}, error) {
	if len(r.data) > 0 {
		return r.data[0], nil
	}

	return nil, errors.New("calling First() but no record found")
}
