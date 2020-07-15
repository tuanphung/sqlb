package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawChain(t *testing.T) {
	statement := RawStatement{
		Raw: "EXPLAIN",
	}

	chain := RawChain([]Sqlizer{statement})

	sql, args, err := chain.ToSql()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestRawChainWithArgs(t *testing.T) {
	statement := RawStatement{
		Raw:  "EXPLAIN",
		Args: []interface{}{"1", 1},
	}

	chain := RawChain([]Sqlizer{statement})

	sql, args, err := chain.ToSql()
	assert.Equal(t, "EXPLAIN", sql, "they should be equal")
	assert.Equal(t, []interface{}{"1", 1}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestRawChainWithSelect(t *testing.T) {
	statement := RawStatement{
		Raw: "EXPLAIN",
	}

	chain := RawChain([]Sqlizer{statement})

	sql, args, err := chain.Select("*").ToSql()
	assert.Equal(t, "EXPLAIN SELECT *", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestRawChainWithFrom(t *testing.T) {
	statement := RawStatement{
		Raw: "EXPLAIN",
	}

	chain := RawChain([]Sqlizer{statement})

	sql, args, err := chain.From("table").ToSql()
	assert.Equal(t, "EXPLAIN FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
