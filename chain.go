package main

import "strings"

type Chain struct {
	statements []Sqlizer
}

func (c Chain) ToSql() (string, []interface{}, error) {
	parts := []string{}
	aggregatedArgs := []interface{}{}

	for _, s := range c.statements {
		sql, args, _ := s.ToSql()
		if sql != "" {
			parts = append(parts, sql)
			aggregatedArgs = append(aggregatedArgs, args...)
		}
	}

	if len(parts) == 0 {
		return "", aggregatedArgs, nil
	}

	if len(parts) == 1 {
		return parts[0], aggregatedArgs, nil
	}

	return strings.Join(parts, " "), aggregatedArgs, nil
}

type ChainBuilder struct {
	Chain Chain
}

func (b ChainBuilder) Raw(raw string, args ...interface{}) RawChain {
	statement := RawStatement{
		Raw:  raw,
		Args: args,
	}

	chain := b.Chain
	chain.statements = append(chain.statements, statement)
	return RawChain(chain)
}

func (b ChainBuilder) Select(columns ...string) SelectChain {
	statement := SelectStatement{
		Columns: columns,
	}

	chain := b.Chain
	chain.statements = append(chain.statements, statement)
	return SelectChain(chain)
}

func (b ChainBuilder) From(tables ...string) FromChain {
	statement := FromStatement{
		Tables: tables,
	}

	chain := b.Chain
	chain.statements = append(chain.statements, statement)
	return FromChain(chain)
}