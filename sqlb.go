package sqlb

func Raw(value string) RawChain {
	return ChainBuilder{}.Raw(value)
}

func Select(columns ...string) SelectChain {
	return ChainBuilder{}.Select(columns...)
}
