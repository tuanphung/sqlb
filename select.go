package sqlb

type SelectChain Chain

func (c SelectChain) ToSql() (string, []interface{}, error) {
	return Chain(c).ToSql()
}

func (c SelectChain) Raw(raw string, args ...interface{}) RawChain {
	return ChainBuilder{Chain(c)}.Raw(raw, args...)
}

func (c SelectChain) From(tables ...string) FromChain {
	return ChainBuilder{Chain(c)}.From(tables...)
}