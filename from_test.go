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
