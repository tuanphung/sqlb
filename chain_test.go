package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainBuilderWithEmpty(t *testing.T) {
	chain := ChainBuilder{}.Chain

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithRaw(t *testing.T) {
	chain := ChainBuilder{}.Raw("EXPLAIN")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithRawAndArgs(t *testing.T) {
	chain := ChainBuilder{}.Raw("EXPLAIN", "a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, []interface{}{"a", "b"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithSelect(t *testing.T) {
	chain := ChainBuilder{}.Select("a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "SELECT a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestChainBuilderWithFrom(t *testing.T) {
	chain := ChainBuilder{}.From("a", "b")

	sql, args, err := chain.ToExpr()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
