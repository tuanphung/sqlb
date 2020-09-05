package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRebindSQL(t *testing.T) {
	sql := "SELECT * FROM table where foo = ? AND bar = ?"

	SetPlaceholder(Dollar)
	assert.Equal(t, "SELECT * FROM table where foo = $1 AND bar = $2", Rebind(sql), "they should be equal")

	SetPlaceholder(Question)
}

func TestChainBuilderWithEmpty(t *testing.T) {
	chain := ChainBuilder{}.Chain

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithRaw(t *testing.T) {
	chain := ChainBuilder{}.Raw("EXPLAIN")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithRawAndArgs(t *testing.T) {
	chain := ChainBuilder{}.Raw("EXPLAIN", "a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, []interface{}{"a", "b"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithSelect(t *testing.T) {
	chain := ChainBuilder{}.Select("a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "SELECT a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithFrom(t *testing.T) {
	chain := ChainBuilder{}.From("a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithOffset(t *testing.T) {
	chain := ChainBuilder{}.Offset(0)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "OFFSET 0", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithLimit(t *testing.T) {
	chain := ChainBuilder{}.Limit(10)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "LIMIT 10", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithOrderBy(t *testing.T) {
	chain := ChainBuilder{}.OrderBy(Order{"foo", false})

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "ORDER BY foo ASC", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

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

func TestAppendChain(t *testing.T) {
	chain1 := From("table").Where(Eq{"foo", "bar"})

	chain2 := Select("*")

	finalChain := chain2.Append(chain1)
	sql, args, _ := finalChain.ToExpr()
	assert.Equal(t, "SELECT * FROM table WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}
