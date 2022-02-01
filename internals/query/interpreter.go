package query

import (
	"fmt"
	"strings"
)

type limitResult struct {
	limit int
	ok    bool
}

type SelectSchema struct {
	TableStr   string
	SelectStr  string
	WhereStr   string
	GroupByStr string
	LimitStr   limitResult
	OrderByStr string
	HavingStr  string
	Bindings   []interface{}
}

func Interpreter() *SelectSchema {
	// var binds []interface{}

	s := &SelectSchema{
		TableStr:   "",
		SelectStr:  "*",
		WhereStr:   "",
		GroupByStr: "",
		LimitStr:   limitResult{0, false},
		OrderByStr: "",
		HavingStr:  "",
		// Bindings:   binds,
	}

	return s
}

func (s *SelectSchema) AppendBinding(bind interface{}) *SelectSchema {
	s.Bindings = append(s.Bindings, bind)

	return s
}

func (s *SelectSchema) Table(table string) *SelectSchema {
	s.TableStr = table

	return s
}

func (s *SelectSchema) Select(cols ...string) *SelectSchema {
	s.SelectStr = strings.Join(cols, ", ")

	return s
}

func (s *SelectSchema) SelectRaw(cols string) *SelectSchema {
	s.SelectStr = cols

	return s
}

func (s *SelectSchema) Where(stmt string, bind interface{}) *SelectSchema {
	if len(s.WhereStr) == 0 {
		s.WhereStr = stmt
		s.AppendBinding(bind)
	} else {
		s.WhereStr = fmt.Sprintf("%s and %s", s.WhereStr, stmt)
		s.AppendBinding(bind)
	}

	return s
}

func (s *SelectSchema) OrWhere(stmt string, bind interface{}) *SelectSchema {
	if len(s.WhereStr) == 0 {
		s.Where(stmt, bind)
	} else {
		s.WhereStr = fmt.Sprintf("%s or %s", s.WhereStr, stmt)
		s.AppendBinding(bind)
	}

	return s
}

func (s *SelectSchema) AndWhere(stmt string, bind interface{}) *SelectSchema {
	s.Where(stmt, bind)

	return s
}

func (s *SelectSchema) WhereRaw(stmt string) *SelectSchema {
	s.WhereStr = stmt

	return s
}

func (s *SelectSchema) OrderBy(col string, sort string) *SelectSchema {
	s.OrderByStr = fmt.Sprintf("%s %s", col, sort)

	return s
}

func (s *SelectSchema) GroupBy(col string) *SelectSchema {
	s.GroupByStr = col

	return s
}

func (s *SelectSchema) Having(stmt string, bind interface{}) *SelectSchema {
	if len(s.HavingStr) == 0 {
		s.HavingStr = stmt
		s.AppendBinding(bind)
	} else {
		s.HavingStr = fmt.Sprintf("%s and %s", s.HavingStr, stmt)
		s.AppendBinding(bind)
	}

	return s
}

func (s *SelectSchema) Limit(lim int) *SelectSchema {
	s.LimitStr = limitResult{lim, true}

	return s
}

func (s *SelectSchema) getLimit() (int, bool) {
	return s.LimitStr.limit, s.LimitStr.ok
}

func (s *SelectSchema) ToSql() string {
	stmt := fmt.Sprintf("select %s from `%s`", s.SelectStr, s.TableStr)

	if len(s.WhereStr) > 0 {
		stmt = fmt.Sprintf("%s where %s", stmt, s.WhereStr)
	}

	if len(s.GroupByStr) > 0 {
		stmt = fmt.Sprintf("%s group by %s", stmt, s.GroupByStr)
	}

	if len(s.HavingStr) > 0 {
		stmt = fmt.Sprintf("%s having %s", stmt, s.HavingStr)
	}

	if len(s.OrderByStr) > 0 {
		stmt = fmt.Sprintf("%s order by %s", stmt, s.OrderByStr)
	}

	if limit, ok := s.getLimit(); ok {
		stmt = fmt.Sprintf("%s limit %d", stmt, limit)
	}

	return stmt
}

func (s *SelectSchema) GetBindings() []interface{} {
	return s.Bindings
}
