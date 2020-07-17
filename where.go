package sqlb

type WhereChain Chain

func (c WhereChain) ToExpr() (string, []interface{}, error) {
	return Chain(c).ToExpr()
}
