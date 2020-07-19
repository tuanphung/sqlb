package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromChain(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	chain := FromChain([]Statement{statement})

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestFromChainWithLimit(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	chain := FromChain([]Statement{statement}).Limit(10)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "FROM a, b LIMIT 10", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestFromChainWithOffset(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	chain := FromChain([]Statement{statement}).Offset(0)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "FROM a, b OFFSET 0", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
