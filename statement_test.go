package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawStatement(t *testing.T) {
	statement := RawStatement{
		Raw: "EXPLAIN",
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestRawStatementWithArgs(t *testing.T) {
	statement := RawStatement{
		Raw:  "EXPLAIN",
		Args: []interface{}{"1", 1},
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, []interface{}{"1", 1}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestSelectStatement(t *testing.T) {
	statement := SelectStatement{
		Columns: []string{"a", "b"},
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "SELECT a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestFromStatement(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestWhereStatement(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			Eq{
				Column: "foo",
				Value:  "bar",
			},
		},
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestLimitStatement(t *testing.T) {
	statement := LimitStatement{
		Limit: 10,
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "LIMIT 10", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestOffsetStatement(t *testing.T) {
	statement := OffsetStatement{
		Offset: 0,
	}

	sql, args, err := statement.ToExpr()
	assert.Equal(t, "OFFSET 0", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestOrderByStatement(t *testing.T) {
	tests := []struct {
		column     string
		descending bool
		wantedSQL  string
	}{
		{"foo", false, "ORDER BY foo ASC"},
		{"foo", true, "ORDER BY foo DESC"},
	}

	for _, test := range tests {
		statement := OrderByStatement{
			Orders: []Order{
				Order{
					Column:     test.column,
					Descending: test.descending,
				},
			},
		}

		sql, args, err := statement.ToExpr()
		assert.Equal(t, test.wantedSQL, sql, "they should be equal")
		assert.Equal(t, 0, len(args), "they should be equal")
		assert.Equal(t, nil, err, "they should be equal")
	}
}
