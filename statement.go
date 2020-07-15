package sqlb

import (
	"fmt"
	"strings"
)

type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}

type RawStatement struct {
	Raw  string
	Args []interface{}
}

func (s RawStatement) ToSql() (string, []interface{}, error) {
	return s.Raw, s.Args, nil
}

type SelectStatement struct {
	Columns []string
}

func (s SelectStatement) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "SELECT", strings.Join(s.Columns, ", ")), nil, nil
}

type FromStatement struct {
	Tables []string
}

func (s FromStatement) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "FROM", strings.Join(s.Tables, ", ")), nil, nil
}

type WhereStatement struct {
	Exprs []Expr
}

func (s WhereStatement) ToSql() (string, []interface{}, error) {
	sql, args, err := And(s.Exprs).GetExpr()
	sql = fmt.Sprintf("WHERE %s", sql)
	return sql, args, err
}
