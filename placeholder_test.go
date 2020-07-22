package sqlb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlaceholderTestSuite struct {
	suite.Suite
}

func (suite *PlaceholderTestSuite) TearDownTest() {
	SetPlaceholder(Question)
}

func (suite *PlaceholderTestSuite) TestSetGetPlaceholder() {
	SetPlaceholder(0)
	assert.Equal(suite.T(), Question, GetPlaceholder(), "they should be equal")

	SetPlaceholder(1)
	assert.Equal(suite.T(), Dollar, GetPlaceholder(), "they should be equal")

	SetPlaceholder(2)
	assert.Equal(suite.T(), Question, GetPlaceholder(), "they should be equal")

	SetPlaceholder(3)
	assert.Equal(suite.T(), Question, GetPlaceholder(), "they should be equal")
}

func (suite *PlaceholderTestSuite) TestStatementWithDollarPlaceholder() {
	SetPlaceholder(Dollar)

	sql, _, _ := Raw("EXPLAIN").Select("id", "name", "abc").From("table").Where(
		Eq{"foo", "bar"},
		Eq{"a", "b"},
	).Offset(0).Limit(10).ToExpr()
	assert.Equal(suite.T(), "EXPLAIN SELECT id, name, abc FROM table WHERE (foo = $1 AND a = $2) OFFSET 0 LIMIT 10", sql, "they should be equal")
}

func TestPlaceholderSuite(t *testing.T) {
	suite.Run(t, new(PlaceholderTestSuite))
}
