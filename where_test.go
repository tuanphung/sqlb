package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicWhere(t *testing.T) {
	query1 := &BasicWhere{
		Column:   "foo",
		Operator: Equal,
		Value:    "bar 1",
	}
	sql, args, _ := query1.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar 1"}, args, "they should be equal")
}

func TestEqWhere(t *testing.T) {
	query := Eq{"foo", "bar"}

	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestAndWhere(t *testing.T) {
	query := And{Eq{"foo", "bar"}}

	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}
