package sqlb

type RawChain Chain

func (c RawChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c RawChain) Raw(raw string) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw)
}

func (c RawChain) Select(columns ...string) SelectChain {
	return ChainBuilder{Chain(c)}.Select(columns...)
}

func (c RawChain) From(tables ...string) FromChain {
	return ChainBuilder{Chain(c)}.From(tables...)
}

func (c RawChain) Where(parts ...Expr) WhereChain {
	return ChainBuilder{Chain(c)}.Where(parts...)
}

func (c RawChain) Limit(limit int64) LimitChain {
	return ChainBuilder{Chain(c)}.Limit(limit)
}

func (c RawChain) Offset(offset int64) OffsetChain {
	return ChainBuilder{Chain(c)}.Offset(offset)
}
