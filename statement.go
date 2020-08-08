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
	if len(s.Args) == 0 {
		return s.Raw, nil, nil
	}

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

type LimitStatement struct {
	Limit int64
}

func (s LimitStatement) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("LIMIT %d", s.Limit), nil, nil
}

type OffsetStatement struct {
	Offset int64
}

func (s OffsetStatement) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("OFFSET %d", s.Offset), nil, nil
}
