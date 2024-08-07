package synchro

import (
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"github.com/zeebo/assert"
	"testing"
)

func TestGenerateExpressionQuery(t *testing.T) {
	expr, err := syntax.ParseExpression(`id in ['foo'] and id >= 'bar' and id < 'baz'`)
	assert.NoError(t, err)

	sql, _, err := db.GenerateExpressionQuery(expr)
	assert.NoError(t, err)
	assert.NotNil(t, sql)
}

func TestGenerateExpressionQuery2(t *testing.T) {
	expr, err := syntax.ParseExpression(`id in 'b'..='z'`)
	assert.NoError(t, err)

	sql, _, err := db.GenerateExpressionQuery(expr)
	assert.NoError(t, err)
	assert.NotNil(t, sql)
}
