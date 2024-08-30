package synchro

import (
	jsoniter "github.com/json-iterator/go"
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

type Foo struct {
	Bar  int    `json:"bar"`
	Bazz string `json:"bazz"`
}

func TestJson(t *testing.T) {
	foo := &Foo{
		Bar:  0,
		Bazz: "",
	}

	str, _ := jsoniter.MarshalToString(foo)
	assert.True(t, len(str) > 0)
}
