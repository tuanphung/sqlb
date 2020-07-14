package sqlb

type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}
