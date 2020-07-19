package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffsetChain(t *testing.T) {
	statement := OffsetStatement{
		Offset: 0,
	}

	chain := OffsetChain([]Statement{statement})

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "OFFSET 0", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestOffsetChainWithLimit(t *testing.T) {
	statement := OffsetStatement{
		Offset: 0,
	}

	chain := OffsetChain([]Statement{statement}).Limit(10)

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "OFFSET 0 LIMIT 10", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
