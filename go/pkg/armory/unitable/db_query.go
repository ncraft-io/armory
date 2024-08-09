package unitable

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
)

func (x *DbQuery) Example() *core.Object {
	if x != nil {
		obj := core.NewObject()
		for _, col := range x.Columns {
			obj.SetValue(strcase.ToLowerCamel(col.Name), col.Example)
		}
		return obj
	}

	return nil
}
