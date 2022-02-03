package query

import (
	"testing"

	"github.com/golang-module/carbon"
)

func TestSelect(t *testing.T) {
	got := Interpreter().Select("*", "`custom_field`").Table("users").ToSql()
	want := "select *, `custom_field` from `users`"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSelectRaw(t *testing.T) {
	got := Interpreter().SelectRaw("name, age, location").Table("users").ToSql()
	want := "select name, age, location from `users`"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWhere(t *testing.T) {
	got := Interpreter().Table("users").Where("id = ?", 1).ToSql()
	want := "select * from `users` where id = ?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAndWhere(t *testing.T) {
	got := Interpreter().Table("users").Where("id = ?", 1).AndWhere("name = ?", "Daison").ToSql()
	want := "select * from `users` where id = ? and name = ?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestOrWhere(t *testing.T) {
	got := Interpreter().Table("users").Where("id = ?", 1).OrWhere("name = ?", "Daison").ToSql()
	want := "select * from `users` where id = ? or name = ?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestOrWhereWithNoInitialClause(t *testing.T) {
	got := Interpreter().Table("users").OrWhere("id = ?", 1).OrWhere("name = ?", "Daison").ToSql()
	want := "select * from `users` where id = ? or name = ?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultipleWhere(t *testing.T) {
	now := carbon.Now().ToDateTimeString()
	got := Interpreter().Table("users").Where("id = ?", 1).Where("name = ?", "Daison").Where("deleted_at >= ?", now).ToSql()
	want := "select * from `users` where id = ? and name = ? and deleted_at >= ?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestWhereRaw(t *testing.T) {
	got := Interpreter().Table("users").WhereRaw("id = 1").ToSql()
	want := "select * from `users` where id = 1"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestOthers(t *testing.T) {
	got := Interpreter().
		Table("users").
		Having("fname like ?", "John%").
		Having("lname like ?", "Doe%").
		OrderBy("created_at", "desc").
		GroupBy("age").
		Limit(100).
		Offset(10).
		ToSql()

	want := "select * from `users` group by age having fname like ? and lname like ? order by created_at desc limit 100 offset 10"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
