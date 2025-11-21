package synchro

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/zeebo/assert"
	"testing"
)

func TestDynamicStruct_NewOf(t *testing.T) {
	ds := NewDynamicStruct(&unitable.Table{
		Id:   "test",
		Name: "test",
		Columns: []*unitable.Column{{
			Name: "id",
			Type: "string",
		}, {
			Name:     "values",
			Type:     "string",
			Repeated: true,
		}},
	})

	obj, _ := core.NewObjectFromKeyValues("id", "1", "values", []string{"v1", "v2"})
	s, _ := ds.NewOf(obj)

	po, _ := ParseObject(s)
	assert.NotNil(t, po)
	vs := po.GetStringArray("values")
	assert.Equal(t, 2, len(vs))
	assert.Equal(t, "v1", vs[0])
}
