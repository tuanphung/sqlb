package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicWhereChain(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			Eq{
				Column: "foo",
				Value:  "bar",
			},
		},
	}

	chain := WhereChain([]Statement{statement})
	expr, args, err := chain.ToExpr()
	assert.Equal(t, "WHERE foo = ?", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestWhereChainWithOffset(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			Eq{
				Column: "foo",
				Value:  "bar",
			},
		},
	}

	chain := WhereChain([]Statement{statement}).Offset(0)
	expr, args, err := chain.ToExpr()
	assert.Equal(t, "WHERE foo = ? OFFSET 0", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestWhereChainWithLimit(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			Eq{
				Column: "foo",
				Value:  "bar",
			},
		},
	}

	chain := WhereChain([]Statement{statement}).Limit(10)
	expr, args, err := chain.ToExpr()
	assert.Equal(t, "WHERE foo = ? LIMIT 10", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestWhereChainWithRaw(t *testing.T) {
	statement := WhereStatement{
		Exprs: []Expr{
			Eq{
				Column: "foo",
				Value:  "bar",
			},
		},
	}

	chain := WhereChain([]Statement{statement}).Raw("ORDER BY a")
	expr, args, err := chain.ToExpr()
	assert.Equal(t, "WHERE foo = ? ORDER BY a", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
