package users

import (
	"fmt"

	"github.com/daison12006013/gorvel/databases"
	"github.com/daison12006013/gorvel/pkg/array"
	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/paginate/searchable"

	sq "github.com/Masterminds/squirrel"
)

func Lists(st *searchable.Table) error {
	db := databases.Resolve()

	// fetch counts
	var total int
	countStmt, countArgs, err := st.QueryCount(Table).ToSql()
	if errors.Handler("error fetching count", err) {
		panic(err)
	}
	db.Raw(countStmt, countArgs...).Scan(&total)

	// fetch the data
	var records []Model
	selectStmt, selectArgs, err := st.QuerySelect(Table).ToSql()
	if errors.Handler("error fetching data", err) {
		panic(err)
	}
	db.Raw(selectStmt, selectArgs...).Scan(&records)

	// reload the pagination data
	st.Paginate.Reconstruct(&records, total)
	return nil
}

func Exists(id *string) (bool, error) {
	if id == nil {
		return false, fmt.Errorf("id should not be null")
	}

	db := databases.Resolve()

	stmt, args, _ := sq.Select("1").From(Table).Where(sq.Eq{"id": &id}).ToSql()
	stmt, args, _ = sq.Expr("select exists("+stmt+") as found", args).ToSql()

	var found bool
	db.Raw(stmt, args).Scan(&found)
	return found, nil
}

// ---

type Finder struct {
	Record *Model
}

func Find(id *string) *Finder {
	db := databases.Resolve()
	user := &Model{}
	db.First(user, id)
	return &Finder{Record: user}
}

func (f *Finder) Delete() bool {
	db := databases.Resolve()
	db.Delete(f.Record)
	return true
}

func (f *Finder) Updates(inputs map[string]interface{}) {
	db := databases.Resolve()

	// only filter updatable fields!
	for k := range inputs {
		if array.In(k, Updatables) < 0 {
			delete(inputs, k)
		}
	}

	db.Model(f.Record).Updates(inputs)
}
