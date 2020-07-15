package sqlb

type RawChain Chain

func (c RawChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}

func (c RawChain) Select(columns ...string) SelectChain {
	return ChainBuilder{Chain(c)}.Select(columns...)
}

func (c RawChain) From(tables ...string) FromChain {
	return ChainBuilder{Chain(c)}.From(tables...)
}
