package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderByChain(t *testing.T) {
	tests := []struct {
		column     string
		descending bool
		wantedSQL  string
	}{
		{"foo", false, "ORDER BY foo ASC"},
		{"foo", true, "ORDER BY foo DESC"},
	}

	for _, test := range tests {
		statement := OrderByStatement{
			Orders: []Order{
				Order{
					Column:     test.column,
					Descending: test.descending,
				},
			},
		}
		chain := OrderByChain([]Statement{statement})

		sql, args, err := chain.ToExpr()
		assert.Equal(t, test.wantedSQL, sql, "they should be equal")
		assert.Equal(t, 0, len(args), "they should be equal")
		assert.Equal(t, nil, err, "they should be equal")
	}
}
