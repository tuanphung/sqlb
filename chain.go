package sqlb

import "strings"

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
