package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitChain(t *testing.T) {
	statement := LimitStatement{
		Limit: 10,
	}

	chain := LimitChain([]Statement{statement})

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "LIMIT 10", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestLimitChainWithOffset(t *testing.T) {
	statement := LimitStatement{
		Limit: 10,
	}

	chain := LimitChain([]Statement{statement}).Offset(0)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "LIMIT 10 OFFSET 0", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
