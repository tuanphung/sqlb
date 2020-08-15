package sqlb

type LimitChain Chain

func (c LimitChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c LimitChain) Offset(offset int64) OffsetChain {
	return ChainBuilder{Chain(c)}.Offset(offset)
}

func (c LimitChain) OrderBy(orders ...Order) OrderByChain {
	return ChainBuilder{Chain(c)}.OrderBy(orders...)
}
