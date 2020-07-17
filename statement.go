package sqlb

import (
	"fmt"
	"strings"
)

type Statement interface {
	ToExpr() (string, []interface{}, error)
}

type RawStatement struct {
	Raw  string
	Args []interface{}
}

func (s RawStatement) ToExpr() (string, []interface{}, error) {
	return s.Raw, s.Args, nil
}

type SelectStatement struct {
	Columns []string
}

func (s SelectStatement) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "SELECT", strings.Join(s.Columns, ", ")), nil, nil
}

type FromStatement struct {
	Tables []string
}

func (s FromStatement) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "FROM", strings.Join(s.Tables, ", ")), nil, nil
}

type WhereStatement struct {
	Exprs []Expr
}

func (s WhereStatement) ToExpr() (string, []interface{}, error) {
	sql, args, err := And(s.Exprs).ToExpr()
	sql = fmt.Sprintf("WHERE %s", sql)
	return sql, args, err
}
