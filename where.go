package sqlb

import (
	"errors"
	"fmt"
	"strings"
)

type WhereChain Chain

func (c WhereChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
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

type Expr interface {
	GetExpr() (string, []interface{}, error)
}

type SingleExpr struct {
	Column   string
	Operator Operator
	Value    interface{}
}

func (o SingleExpr) GetExpr() (string, []interface{}, error) {
	if o.Operator == "IN" {
		if array, ok := o.Value.([]interface{}); ok {
			variables := []string{}
			for i := 0; i < len(array); i++ {
				variables = append(variables, "?")
			}

			return fmt.Sprintf("%s IN (%s)", o.Column, strings.Join(variables, ",")), array, nil
		}
		return "", nil, errors.New("error GetExpr IN: value is not []interface{}")
	} else {
		return fmt.Sprintf("%s %s ?", o.Column, o.Operator), []interface{}{o.Value}, nil
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

func (o Conjunction) GetExpr() (string, []interface{}, error) {
	statements := []string{}
	finalArgs := []interface{}{}

	for _, query := range o.Exprs {
		statement, args, _ := query.GetExpr()

		statements = append(statements, statement)
		finalArgs = append(finalArgs, args...)
	}

	format := "(%s)"
	if len(statements) == 1 {
		format = "%s"
	}

	return fmt.Sprintf(format, strings.Join(statements, fmt.Sprintf(" %s ", o.Operator))), finalArgs, nil
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
