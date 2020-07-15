package sqlb

import (
	"errors"
	"fmt"
	"strings"
)

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

type Expr interface {
	GetExpr() (string, []interface{}, error)
}

type SingleExpr struct {
	Column   string
	Operator Operator
	Value    interface{}
}

func (e SingleExpr) GetExpr() (string, []interface{}, error) {
	if e.Operator == "IN" {
		if array, ok := e.Value.([]interface{}); ok {
			placeholders := []string{}
			for i := 0; i < len(array); i++ {
				placeholders = append(placeholders, "?")
			}

			return fmt.Sprintf("%s IN (%s)", e.Column, strings.Join(placeholders, ",")), array, nil
		}
		return "", nil, errors.New("error GetExpr IN: value is not []interface{}")
	} else {
		return fmt.Sprintf("%s %s ?", e.Column, e.Operator), []interface{}{e.Value}, nil
	}
}

type ConjunctionOperator string

const (
	AndOperator ConjunctionOperator = "AND"
	OrOperator  ConjunctionOperator = "OR"
)

type Conjunction struct {
	Operator ConjunctionOperator
	Exprs    []Expr
}

func (c Conjunction) GetExpr() (string, []interface{}, error) {
	if len(c.Exprs) == 0 {
		return "", nil, nil
	}

	if len(c.Exprs) == 1 {
		return c.Exprs[0].GetExpr()
	}

	exprs := []string{}
	aggregatedArgs := []interface{}{}

	for _, e := range c.Exprs {
		expr, args, _ := e.GetExpr()

		exprs = append(exprs, expr)
		aggregatedArgs = append(aggregatedArgs, args...)
	}

	return fmt.Sprintf("(%s)", strings.Join(exprs, fmt.Sprintf(" %s ", c.Operator))), aggregatedArgs, nil
}

type Eq struct {
	Column string
	Value  interface{}
}

func (o Eq) GetExpr() (string, []interface{}, error) {
	return SingleExpr{o.Column, EqualOperator, o.Value}.GetExpr()
}

type NotEq struct {
	Column string
	Value  interface{}
}

func (o NotEq) GetExpr() (string, []interface{}, error) {
	return SingleExpr{o.Column, NotEqualOperator, o.Value}.GetExpr()
}

type In struct {
	Column string
	Value  []interface{}
}

func (o In) GetExpr() (string, []interface{}, error) {
	return SingleExpr{o.Column, InOperator, o.Value}.GetExpr()
}

type And []Expr

func (o And) GetExpr() (string, []interface{}, error) {
	conj := Conjunction{AndOperator, o}
	return conj.GetExpr()
}

type Or []Expr

func (o Or) GetExpr() (string, []interface{}, error) {
	conj := Conjunction{OrOperator, o}
	return conj.GetExpr()
}
