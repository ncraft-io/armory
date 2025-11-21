package synchro

import "github.com/ncraft-io/armory/go/pkg/armory/unitable"

func FilterOutId(maps map[string]interface{}, table *unitable.Table) map[string]interface{} {
	out := make(map[string]interface{})
	index := table.ColumnIndex()
	for k, v := range maps {
		if k == "id" {
			continue
		}

		if col, ok := index[k]; ok && col.IsStringField() {
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
