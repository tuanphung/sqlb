package sqlb

type WhereChain Chain

func (c WhereChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}
