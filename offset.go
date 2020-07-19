package sqlb

type OffsetChain Chain

func (c OffsetChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c OffsetChain) Limit(limit int64) LimitChain {
	return ChainBuilder{Chain(c)}.Limit(limit)
}
