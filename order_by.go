package sqlb

type OrderByChain Chain

func (c OrderByChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}

func (c OrderByChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}
