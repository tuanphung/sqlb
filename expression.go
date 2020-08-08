package sqlb

import (
	"fmt"
	"strings"
)

type Expr interface {
	ToExpr() (string, []interface{}, error)
}

type Operator string

const (
	EqualOperator          Operator = "="
	NotEqualOperator       Operator = "<>"
	GreaterOperator        Operator = ">"
	GreaterOrEqualOperator Operator = ">="
	LessOperator           Operator = "<"
	LessOrEqualOperator    Operator = "<="
	InOperator             Operator = "IN"
	LikeOperator           Operator = "LIKE"
)

type Nested struct {
	Column   string
	Operator Operator
	Value    Expr
}

func (e Nested) ToExpr() (string, []interface{}, error) {
	expr, args, _ := e.Value.ToExpr()
	return fmt.Sprintf("%s %s (%s)", e.Column, e.Operator, expr), args, nil
}

type Eq struct {
	Column string
	Value  interface{}
}

func (e Eq) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("%s = ?", e.Column), []interface{}{e.Value}, nil
}

type NotEq struct {
	Column string
	Value  interface{}
}

func (e NotEq) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("%s <> ?", e.Column), []interface{}{e.Value}, nil
}

type Like struct {
	Column string
	Value  interface{}
}

func (e Like) ToExpr() (string, []interface{}, error) {
	return fmt.Sprintf("%s LIKE ?", e.Column), []interface{}{e.Value}, nil
}

type In struct {
	Column string
	Value  []interface{}
}

func (e In) ToExpr() (string, []interface{}, error) {
	placeholders := []string{}
	for i := 0; i < len(e.Value); i++ {
		placeholders = append(placeholders, "?")
	}

	return fmt.Sprintf("%s IN (%s)", e.Column, strings.Join(placeholders, ",")), e.Value, nil
}

type Order struct {
	Column     string
	Descending bool
}

func (e Order) ToExpr() (string, []interface{}, error) {
	order := "ASC"
	if e.Descending {
		order = "DESC"
	}

	return fmt.Sprintf("%s %s", e.Column, order), nil, nil
}

type ConjunctionOperator string

const (
	AndOperator ConjunctionOperator = "AND"
	OrOperator  ConjunctionOperator = "OR"
)

type conjunctionExpr struct {
	Operator ConjunctionOperator
	Exprs    []Expr
}

func (c conjunctionExpr) ToExpr() (string, []interface{}, error) {
	if len(c.Exprs) == 0 {
		return "", nil, nil
	}

	if len(c.Exprs) == 1 {
		return c.Exprs[0].ToExpr()
	}

	exprs := []string{}
	aggregatedArgs := []interface{}{}

	for _, e := range c.Exprs {
		expr, args, _ := e.ToExpr()

		exprs = append(exprs, expr)
		aggregatedArgs = append(aggregatedArgs, args...)
	}

	return fmt.Sprintf("(%s)", strings.Join(exprs, fmt.Sprintf(" %s ", c.Operator))), aggregatedArgs, nil
}

type And []Expr

func (e And) ToExpr() (string, []interface{}, error) {
	conj := conjunctionExpr{AndOperator, e}
	return conj.ToExpr()
}

type Or []Expr

func (e Or) ToExpr() (string, []interface{}, error) {
	conj := conjunctionExpr{OrOperator, e}
	return conj.ToExpr()
}
