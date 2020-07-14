package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	sql, args, _ := Select("id", "name", "abc").ToSql()
	assert.Equal(t, "SELECT id, name, abc", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestSelectFrom(t *testing.T) {
	sql, args, _ := Select("id", "name", "abc").From("table").ToSql()
	assert.Equal(t, "SELECT id, name, abc FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestRawSelectFrom(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").ToSql()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestRawSelectFromWithArgs(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Raw("WHERE foo = ?", "bar").ToSql()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}
