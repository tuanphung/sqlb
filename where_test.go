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
				Value: "bar",
			},
		},
	}

	chain := WhereChain([]Statement{statement})
	expr, args, err := chain.ToExpr()
	assert.Equal(t, "WHERE foo = ?", expr, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
