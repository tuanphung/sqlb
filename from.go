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
