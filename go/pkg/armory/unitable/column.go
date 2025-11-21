package unitable

import (
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
)

func (x *Column) IsTypeValid() bool {
	if x != nil && len(x.Type) > 0 {
		switch x.Type {
		case "bool", "integer", "float", "string":
			return true
		}
	}
	return false
}

func (x *Column) IsStringField() bool {
	return x != nil && x.Type == "string"
}

func (x *Column) ToFieldInfo() query.Field {
	if x != nil {
		field := query.Field{
			Type:     0,
			Repeated: x.Repeated,
		}
		switch x.Type {
		case "bool":
			field.Type = query.FieldTypeBool
		case "integer":
			field.Type = query.FieldTypeInteger
		case "float":
			field.Type = query.FieldTypeFloat
		case "string":
			switch x.Format {
			case "time", "datetime", "timestamp":
				field.Type = query.FieldTypeDatetime
			case "bytes":
				field.Type = query.FieldTypeBytes
			case "geometry":
				field.Type = query.FieldTypeGeometry
			case "json":
				field.Type = query.FieldTypeJSON
			default:
				field.Type = query.FieldTypeString
			}
		}
		return field
	}

	return query.Field{}
}
