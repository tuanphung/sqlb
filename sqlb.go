package sqlb

func Raw(raw string, args ...interface{}) Chain {
	return ChainBuilder{}.Raw(raw, args...)
}

func Select(columns ...string) Chain {
	return ChainBuilder{}.Select(columns...)
}

func From(tables ...string) Chain {
	return ChainBuilder{}.From(tables...)
}

func Where(parts ...Expr) Chain {
	return ChainBuilder{}.Where(parts...)
}

func Limit(limit int64) Chain {
	return ChainBuilder{}.Limit(limit)
}

func Offset(offset int64) Chain {
	return ChainBuilder{}.Offset(offset)
}

func OrderBy(orders ...Order) Chain {
	return ChainBuilder{}.OrderBy(orders...)
}