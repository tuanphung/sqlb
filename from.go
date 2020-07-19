package sqlb

type FromChain Chain

func (c FromChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c FromChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}

func (c FromChain) Where(parts ...Expr) WhereChain {
	return ChainBuilder{Chain(c)}.Where(parts...)
}

func (c FromChain) Limit(limit int64) LimitChain {
	return ChainBuilder{Chain(c)}.Limit(limit)
}

func (c FromChain) Offset(offset int64) OffsetChain {
	return ChainBuilder{Chain(c)}.Offset(offset)
}
