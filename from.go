package main

import (
	"fmt"
	"strings"
)

type FromStatement struct {
	Tables []string
}

func (s FromStatement) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf("%s %s", "FROM", strings.Join(s.Tables, ", ")), nil, nil
}

type FromChain Chain

func (c FromChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}

func (c FromChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}
