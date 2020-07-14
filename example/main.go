package main

import (
	"fmt"

	sb "github.com/tuanphung/sqlb"
)

func main() {
	sql, args, _ := sb.Select("*").From("table").Raw("WHERE foo = ?", "bar").ToSql()
	fmt.Println(sql)
	fmt.Println(args...)
}
