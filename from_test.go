package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromStatement(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	sql, args, err := statement.ToSql()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestFromChain(t *testing.T) {
	statement := FromStatement{
		Tables: []string{"a", "b"},
	}

	chain := FromChain{
		statements: []Sqlizer{statement},
	}

	sql, args, err := chain.ToSql()
	assert.Equal(t, "FROM a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}