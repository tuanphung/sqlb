package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawStatement(t *testing.T) {
	statement := RawStatement{
		Raw: "EXPLAIN",
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestRawStatementWithArgs(t *testing.T) {
	statement := RawStatement{
		Raw:  "EXPLAIN",
		Args: []interface{}{"1", 1},
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, []interface{}{"1", 1}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestSelectStatement(t *testing.T) {
	statement := SelectStatement{
		Columns: []string{"a", "b"},
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "SELECT a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestFromStatement(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestWhereStatement(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			SingleExpr{
				Column:   "foo",
				Operator: EqualOperator,
				Value:    "bar",
			},
		},
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
