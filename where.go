package sqlb

type WhereChain Chain

func (c WhereChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c WhereChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args)
}

func (c WhereChain) Limit(limit int64) LimitChain {
	return ChainBuilder{Chain(c)}.Limit(limit)
}

func (c WhereChain) Offset(offset int64) OffsetChain {
	return ChainBuilder{Chain(c)}.Offset(offset)
}
