package unitable

import "github.com/mojo-lang/db/go/pkg/mojo/db"

func (x *Column) IsTypeValid() bool {
	if x != nil && len(x.Type) > 0 {
		switch x.Type {
		case "bool", "integer", "float", "string":
			return true
		}
	}
	return false
}

func (x *Column) ToFieldInfo() db.FieldInfo {
	if x != nil {
		info := db.FieldInfo{
			Type:     0,
			Repeated: x.Repeated,
		}
		switch x.Type {
		case "integer":
			info.Type = db.FieldTypeInteger
		case "number":
			info.Type = db.FieldTypeFloat
		case "string":
			switch x.Format {
			case "time", "datetime", "timestamp":
				info.Type = db.FieldTypeDatetime
			case "bytes":
				info.Type = db.FieldTypeBytes
			default:
				info.Type = db.FieldTypeString
			}
		}
		return info
	}

	return db.FieldInfo{}
}
