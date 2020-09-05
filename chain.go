package sqlb

import (
	"strconv"
	"strings"
)

type Chain []Statement

func (c Chain) ToExpr() (string, []interface{}, error) {
	parts := []string{}
	aggregatedArgs := []interface{}{}

	for _, s := range c {
		sql, args, _ := s.ToExpr()
		if sql != "" {
			parts = append(parts, sql)
			aggregatedArgs = append(aggregatedArgs, args...)
		}
	}

	sql := ""
	if len(parts) > 0 {
		sql = strings.Join(parts, " ")
	}

	// Rebind argument placeholder
	sql = rebind(GetPlaceholder(), sql)

	return sql, aggregatedArgs, nil
}

func (c Chain) Append(chain Chain) Chain {
	return ChainBuilder{c}.Append(chain)
}

func (c Chain) Raw(raw string, args ...interface{}) Chain {
	return ChainBuilder{c}.Raw(raw, args...)
}

func (c Chain) Select(columns ...string) Chain {
	return ChainBuilder{c}.Select(columns...)
}

func (c Chain) From(tables ...string) Chain {
	return ChainBuilder{c}.From(tables...)
}

func (c Chain) Where(parts ...Expr) Chain {
	return ChainBuilder{c}.Where(parts...)
}

func (c Chain) Limit(limit int64) Chain {
	return ChainBuilder{c}.Limit(limit)
}

func (c Chain) Offset(offset int64) Chain {
	return ChainBuilder{c}.Offset(offset)
}

func (c Chain) OrderBy(orders ...Order) Chain {
	return ChainBuilder{c}.OrderBy(orders...)
}

func Rebind(sql string) string {
	return rebind(GetPlaceholder(), sql)
}

func rebind(placeholder PlaceholderType, sql string) string {
	switch placeholder {
	case Question:
		return sql
	}

	chunks := []byte{}
	var i, j int

	for i = strings.Index(sql, "?"); i != -1; i = strings.Index(sql, "?") {
		chunks = append(chunks, sql[:i]...)

		switch placeholder {
		case Dollar:
			chunks = append(chunks, '$')
		}

		j++
		chunks = strconv.AppendInt(chunks, int64(j), 10)

		sql = sql[i+1:]
	}

	return string(append(chunks, sql...))
}

type ChainBuilder struct {
	Chain Chain
}

func (b ChainBuilder) Append(c Chain) Chain {
	chain := b.Chain
	chain = append(chain, c...)
	return chain
}

func (b ChainBuilder) Raw(raw string, args ...interface{}) Chain {
	statement := RawStatement{
		Raw:  raw,
		Args: args,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) Select(columns ...string) Chain {
	statement := SelectStatement{
		Columns: columns,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) From(tables ...string) Chain {
	statement := FromStatement{
		Tables: tables,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) Where(exprs ...Expr) Chain {
	statement := WhereStatement{
		Exprs: exprs,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) Offset(offset int64) Chain {
	statement := OffsetStatement{
		Offset: offset,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) Limit(limit int64) Chain {
	statement := LimitStatement{
		Limit: limit,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}

func (b ChainBuilder) OrderBy(orders ...Order) Chain {
	statement := OrderByStatement{
		Orders: orders,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return chain
}
