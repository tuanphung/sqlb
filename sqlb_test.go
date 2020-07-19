package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	sql, args, _ := Select("id", "name", "abc").ToExpr()
	assert.Equal(t, "SELECT id, name, abc", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestSelectFrom(t *testing.T) {
	sql, args, _ := Select("id", "name", "abc").From("table").ToExpr()
	assert.Equal(t, "SELECT id, name, abc FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestRawSelectFrom(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestRawSelectFromWithArgs(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Raw("WHERE foo = ?", "bar").ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestRawSelectFromWhereWithArgs(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(Eq{"foo", "bar"}).ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestRawSelectFromWhereWithArgs2(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(And{Eq{"foo", "bar"}, Eq{"a", "b"}}).ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE (foo = ? AND a = ?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar", "b"}, args, "they should be equal")
}

func TestRawSelectFromWhereWithArgs3(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(Eq{"foo", "bar"}, Eq{"a", "b"}).ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE (foo = ? AND a = ?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar", "b"}, args, "they should be equal")
}

func TestRawSelectFromWhereOrWithArgs2(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(Or{Eq{"foo", "bar"}, Eq{"a", "b"}}).ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE (foo = ? OR a = ?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar", "b"}, args, "they should be equal")
}

func TestUsage1(t *testing.T) {
	sql, args, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(Or{Eq{"foo", "bar"}, Eq{"a", "b"}}).Offset(0).Limit(10).ToExpr()
	assert.Equal(t, "EXPLAIN SELECT id, name, abc FROM table WHERE (foo = ? OR a = ?) OFFSET 0 LIMIT 10", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar", "b"}, args, "they should be equal")
}
