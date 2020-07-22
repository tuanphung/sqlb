package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	e := Eq{"foo", "bar"}

	sql, args, _ := e.ToExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestNotEq(t *testing.T) {
	e := NotEq{"foo", "bar"}

	sql, args, _ := e.ToExpr()
	assert.Equal(t, "foo <> ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestIn(t *testing.T) {
	e := In{"foo", []interface{}{"bar"}}
	sql, args, _ := e.ToExpr()
	assert.Equal(t, "foo IN (?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestAnd(t *testing.T) {
	e := And{Eq{"foo", "bar"}}

	sql, args, _ := e.ToExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestLike(t *testing.T) {
	e := And{Like{"foo", "bar"}}

	sql, args, _ := e.ToExpr()
	assert.Equal(t, "foo LIKE ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestEmptyExpr(t *testing.T) {
	e := And{}

	sql, args, _ := e.ToExpr()
	assert.Equal(t, "", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}

func TestNestedExpr(t *testing.T) {
	e := Nested{"id", InOperator, Select("id").From("table").Where(Eq{"foo", "bar"})}

	expr, args, _ := e.ToExpr()
	assert.Equal(t, "id IN (SELECT id FROM table WHERE foo = ?)", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}
