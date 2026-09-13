package synchro

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"github.com/zeebo/assert"
	"sync"
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

func TestDynamicStructConcurrentInitialization(t *testing.T) {
	ds := NewDynamicStruct(&unitable.Table{Columns: []*unitable.Column{{Name: "id", Type: "string"}, {Name: "enabled", Type: "bool"}}})
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				_ = ds.New()
				_ = ds.NewSliceOf()
			}
		}()
	}
	wg.Wait()
}
