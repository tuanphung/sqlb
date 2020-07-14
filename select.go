package main

import (
	"fmt"
	"strings"
)

type SelectStatement struct {
	Columns []string
}

func (s SelectStatement) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "SELECT", strings.Join(s.Columns, ", ")), nil, nil
}

type SelectChain Chain

func (c SelectChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}

func (c SelectChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}

func (c SelectChain) From(tables ...string) FromChain {
	return ChainBuilder{Chain(c)}.From(tables...)
}
