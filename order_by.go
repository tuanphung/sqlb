package sqlb

type OrderByChain Chain

func (c OrderByChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c OrderByChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}

func (c OrderByChain) Offset(offset int64) OffsetChain {
	return ChainBuilder{Chain(c)}.Offset(offset)
}
