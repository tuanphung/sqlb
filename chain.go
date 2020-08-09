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

func (b ChainBuilder) Raw(raw string, args ...interface{}) RawChain {
	statement := RawStatement{
		Raw:  raw,
		Args: args,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return RawChain(chain)
}

func (b ChainBuilder) Select(columns ...string) SelectChain {
	statement := SelectStatement{
		Columns: columns,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return SelectChain(chain)
}

func (b ChainBuilder) From(tables ...string) FromChain {
	statement := FromStatement{
		Tables: tables,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return FromChain(chain)
}

func (b ChainBuilder) Where(exprs ...Expr) WhereChain {
	statement := WhereStatement{
		Exprs: exprs,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return WhereChain(chain)
}

func (b ChainBuilder) Offset(offset int64) OffsetChain {
	statement := OffsetStatement{
		Offset: offset,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return OffsetChain(chain)
}

func (b ChainBuilder) Limit(limit int64) LimitChain {
	statement := LimitStatement{
		Limit: limit,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return LimitChain(chain)
}

func (b ChainBuilder) OrderBy(orders ...Order) OrderByChain {
	statement := OrderByStatement{
		Orders: orders,
	}

	chain := b.Chain
	chain = append(chain, statement)
	return OrderByChain(chain)
}

// Convenience methods to initialize statement chain
func Raw(value string, args ...interface{}) RawChain {
	return ChainBuilder{}.Raw(value, args...)
}

func Select(columns ...string) SelectChain {
	return ChainBuilder{}.Select(columns...)
}
