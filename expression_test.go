package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	e := Eq{"foo", "bar"}

	sql, args, _ := e.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestNotEq(t *testing.T) {
	e := NotEq{"foo", "bar"}

	sql, args, _ := e.GetExpr()
	assert.Equal(t, "foo <> ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestIn(t *testing.T) {
	e := In{"foo", []interface{}{"bar"}}
	sql, args, _ := e.GetExpr()
	assert.Equal(t, "foo IN (?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestAnd(t *testing.T) {
	e := And{Eq{"foo", "bar"}}

	sql, args, _ := e.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestEmptyExpr(t *testing.T) {
	e := And{}

	sql, args, _ := e.GetExpr()
	assert.Equal(t, "", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
}
