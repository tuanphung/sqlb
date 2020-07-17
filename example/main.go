package main

import (
	"fmt"

	sb "github.com/tuanphung/sqlb"
)

func main() {
	testSQ()
	testSB()

	if false && (false || true) {
		fmt.Println("OK")
	}
}

func testSQ() {
	// users := sq.Select("*").From("users").Where(sq.And{sq.Eq{"a": "b"}, sq.Or{sq.Eq{"c": "d", "e": "f"}, sq.Eq{"g": "h"}}})

	// sql, args, _ := users.ToExpr()

	// fmt.Println(sql)
	// fmt.Println(args...)
}

func testSB() {

	users := sb.Select("*").From("users").Where(
		sb.Eq{"a", "b"},
		sb.Or{
			sb.And{
				sb.Eq{"c", "d"},
				sb.Eq{"e", "f"},
			},
			sb.Eq{"g", "h"},
		},
	)

	sql, args, _ := users.ToExpr()

	fmt.Println(sql)
	fmt.Println(args...)
}
