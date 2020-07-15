package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicWhere(t *testing.T) {
	query1 := &SingleExpr{
		Column:   "foo",
		Operator: EqualOperator,
		Value:    "bar 1",
	}
	sql, args, _ := query1.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar 1"}, args, "they should be equal")

	query2 := &SingleExpr{
		Column:   "foo",
		Operator: InOperator,
		Value:    "bar 1",
	}
	sql2, args2, err2 := query2.GetExpr()
	assert.Equal(t, "", sql2, "they should be equal")
	assert.Equal(t, 0, len(args2), "they should be equal")
	assert.Equal(t, "error GetExpr IN: value is not []interface{}", err2.Error(), "they should be equal")
}

func TestEqWhere(t *testing.T) {
	query := Eq{"foo", "bar"}

	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestNotEqWhere(t *testing.T) {
	query := NotEq{"foo", "bar"}

	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo <> ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestInExpr(t *testing.T) {
	query := In{"foo", []interface{}{"bar"}}
	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo IN (?)", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}

func TestAndWhere(t *testing.T) {
	query := And{Eq{"foo", "bar"}}

	sql, args, _ := query.GetExpr()
	assert.Equal(t, "foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
}
