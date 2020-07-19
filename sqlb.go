package sqlb

func Raw(value string, args ...interface{}) RawChain {
	return ChainBuilder{}.Raw(value, args...)
}

func Select(columns ...string) SelectChain {
	return ChainBuilder{}.Select(columns...)
}
