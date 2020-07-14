package main

type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}
