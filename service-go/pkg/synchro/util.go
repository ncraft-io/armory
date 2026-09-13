package synchro

import (
	"github.com/iancoleman/strcase"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
)

func FilterOutId(maps map[string]interface{}, table *unitable.Table) map[string]interface{} {
	out := make(map[string]interface{})
	index := table.ColumnIndex()
	for k, v := range maps {
		k = strcase.ToSnake(k)
		if k == "id" {
			continue
		}

		if col, ok := index[k]; ok && col.IsStringField() {
			if col.Repeated {
				switch values := v.(type) {
				case []string:
					out[k] = StringArray(values)
					continue
				case []interface{}:
					array := make(StringArray, len(values))
					for i, value := range values {
						array[i], _ = value.(string)
					}
					out[k] = array
					continue
				}
			}
			if val, ok := v.(string); ok && len(val) == 0 {
				switch col.Format {
				case "time", "datetime", "timestamp", "geometry":
					out[k] = nil
					continue
				}
			}
		}

		out[k] = v
	}

	return out
}
