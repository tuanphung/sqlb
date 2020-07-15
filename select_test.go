package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectChain(t *testing.T) {
	statement := SelectStatement{
		Columns: []string{"a", "b"},
	}

	chain := SelectChain([]Sqlizer{statement})

	sql, args, err := chain.ToSql()
	assert.Equal(t, "SELECT a, b", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestSelectChainWithRaw(t *testing.T) {
	statement := SelectStatement{
		Columns: []string{"a", "b"},
	}

	chain := SelectChain([]Sqlizer{statement})

	sql, args, err := chain.Raw("FROM table WHERE foo = ?", "bar").ToSql()
	assert.Equal(t, "SELECT a, b FROM table WHERE foo = ?", sql, "they should be equal")
	assert.Equal(t, []interface{}{"bar"}, args, "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}

func TestSelectChainWithFrom(t *testing.T) {
	statement := SelectStatement{
		Columns: []string{"a", "b"},
	}

	chain := SelectChain([]Sqlizer{statement})

	sql, args, err := chain.From("table").ToSql()
	assert.Equal(t, "SELECT a, b FROM table", sql, "they should be equal")
	assert.Equal(t, 0, len(args), "they should be equal")
	assert.Equal(t, nil, err, "they should be equal")
}
