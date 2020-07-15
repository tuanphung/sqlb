package sqlb

import (
	"errors"
	"fmt"
	"strings"
)

type WhereStatement struct {
	Parts []WherePart
}

func (s WhereStatement) ToSql() (string, []interface{}, error) {
	sql, args, err := And(s.Parts).GetExpr()
	sql = fmt.Sprintf("WHERE %s", sql)
	return sql, args, err
}

type WhereChain Chain

func (c WhereChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}

type Operator string

const (
	Equal          Operator = "="
	NotEqual       Operator = "<>"
	Greater        Operator = ">"
	GreaterOrEqual Operator = ">="
	Less           Operator = "<"
	LessOrEqual    Operator = "<="
	In             Operator = "IN"
	Like           Operator = "LIKE"
)

type WherePart interface {
	GetExpr() (string, []interface{}, error)
}

type BasicWhere struct {
	Column   string
	Operator Operator
	Value    interface{}
}

func (o BasicWhere) GetExpr() (string, []interface{}, error) {
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
	andOperator ConjunctionOperator = "AND"
	orOperator  ConjunctionOperator = "OR"
)

type ConjunctionWhere struct {
	Operator ConjunctionOperator
	Parts    []WherePart
}

func (o ConjunctionWhere) GetExpr() (string, []interface{}, error) {
	statements := []string{}
	finalArgs := []interface{}{}

	for _, query := range o.Parts {
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

func AndWhere(parts []WherePart) ConjunctionWhere {
	return ConjunctionWhere{andOperator, parts}
}

func OrWhere(parts []WherePart) ConjunctionWhere {
	return ConjunctionWhere{orOperator, parts}
}

type Eq struct {
	Column string
	Value  interface{}
}

func (o Eq) GetExpr() (string, []interface{}, error) {
	return BasicWhere{o.Column, Equal, o.Value}.GetExpr()
}

type NotEq struct {
	Column string
	Value  interface{}
}

func (o NotEq) GetExpr() (string, []interface{}, error) {
	return BasicWhere{o.Column, Equal, o.Value}.GetExpr()
}

type And []WherePart

func (o And) GetExpr() (string, []interface{}, error) {
	conj := AndWhere(o)
	return conj.GetExpr()
}

type Or []WherePart

func (o Or) GetExpr() (string, []interface{}, error) {
	conj := OrWhere(o)
	return conj.GetExpr()
}
