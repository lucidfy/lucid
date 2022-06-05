package users

import (
	"fmt"
	"net/http"

	"github.com/golang-module/carbon"
	"github.com/lucidfy/lucid/databases"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/hash"
	"github.com/lucidfy/lucid/pkg/functions/php"
	"github.com/lucidfy/lucid/pkg/searchable"

	sq "github.com/Masterminds/squirrel"
)

func Lists(st *searchable.Table) *errors.AppError {
	db := databases.Resolve()
	defer databases.Close(db)

	// fetch counts
	var total int
	countStmt, countArgs, err := st.QueryCount(Table).ToSql()
	if err != nil {
		return errors.InternalServerError("st.QueryCount() error", err)
	}
	db.Raw(countStmt, countArgs...).Scan(&total)

	// fetch the data
	var records []Model
	selectStmt, selectArgs, err := st.QuerySelect(Table).ToSql()
	if err != nil {
		return errors.InternalServerError("st.QuerySelect() error", err)
	}
	db.Raw(selectStmt, selectArgs...).Scan(&records)

	// reload the pagination data
	st.Paginate.Reconstruct(&records, total)
	return nil
}

func Exists(column string, value *string) *errors.AppError {
	if value == nil {
		return &errors.AppError{
			Error:   fmt.Errorf("app.models.users.model@Exists: Value should not be null"),
			Message: "Value should not be null",
			Code:    http.StatusInternalServerError,
		}
	}

	db := databases.Resolve()
	defer databases.Close(db)

	stmt, args, _ := sq.Select("1").From(Table).Where(sq.Eq{column: *value}).ToSql()
	stmt, args, _ = sq.Expr("select exists("+stmt+") as found", args).ToSql()

	var found bool
	db.Raw(stmt, args).Scan(&found)
	if !found {
		return &errors.AppError{
			Error:   fmt.Errorf("app.models.users.model@Exists: (%s: %s) record not found", column, *value),
			Message: "Record not found",
			Code:    http.StatusNotFound,
		}
	}

	return nil
}

func Create(i interface{}) (*Finder, *errors.AppError) {
	db := databases.Resolve()
	defer databases.Close(db)

	// here, we call the sanitizer function
	i, app_err := sanitize(i)
	inputs := i.(map[string]interface{})

	if app_err != nil {
		return nil, app_err
	}

	// here, we validate if the email existence is present
	// we don't wan't to flood our database just by throwing
	// an error of duplicate entry check.
	record := &Model{}
	email := inputs["email"].(string)

	// finding the record should be null
	if exists := Exists("email", &email); exists == nil {
		return nil, &errors.AppError{
			Error:   fmt.Errorf("app.models.users.model@Create: email %s already exists", email),
			Message: fmt.Sprintf("Email %s already exist", email),
			Code:    http.StatusInternalServerError,
		}
	}

	// pre-fill the created_at and updated_at
	now := carbon.Now().ToDateTimeString()
	inputs["created_at"] = now
	inputs["updated_at"] = now

	// create the record, then check if there are error
	err := db.Model(&record).Create(inputs).Error
	if err != nil {
		return nil, &errors.AppError{
			Error:   err,
			Message: "Gorm Error",
			Code:    http.StatusInternalServerError,
		}
	}

	// now return it with the struct finder
	return &Finder{Model: record}, nil
}

// ---

type Finder struct {
	Model *Model
}

func Find(id interface{}, col interface{}) (*Finder, *errors.AppError) {
	db := databases.Resolve()
	defer databases.Close(db)
	record := new(Model)

	if col == nil || col.(string) == "" {
		col = "id"
	}

	err := db.First(record, col.(string)+"=?", id).Error
	if err != nil {
		return nil, &errors.AppError{
			Error:   err,
			Message: "Gorm Error",
			Code:    http.StatusInternalServerError,
		}
	}

	return &Finder{Model: record}, nil
}

func (f *Finder) Updates(i interface{}) *errors.AppError {
	db := databases.Resolve()
	defer databases.Close(db)

	// here, we call the sanitizer function
	i, app_err := sanitize(i)
	inputs := i.(map[string]interface{})

	if app_err != nil {
		return app_err
	}

	// here, we can safely update the inputs
	db.Model(f.Model).Updates(inputs)
	return nil
}

func (f *Finder) Delete() bool {
	db := databases.Resolve()
	defer databases.Close(db)
	db.Delete(f.Model)
	return true
}

func sanitize(i interface{}) (interface{}, *errors.AppError) {
	inputs := i.(map[string]interface{})

	// only filter updatable fields!
	for k := range inputs {
		if php.InArray(k, Updatables) < 0 {
			delete(inputs, k)
		}
	}

	// here, we check if password is present
	// then we need to hash the password
	if pw, ok := inputs["password"]; ok {
		password := pw.(string)
		if len(password) > 0 {
			hashed, app_err := hash.Make(password)
			inputs["password"] = hashed
			if app_err != nil {
				return inputs, app_err
			}
		}
	}

	return inputs, nil
}
