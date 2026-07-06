package synchro

import (
	"fmt"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"reflect"
	"strings"
	"time"
)

const LowerCamel = "lowerCamel"

type DynamicStruct struct {
	Fields    []reflect.StructField
	Type      reflect.Type
	JsonStyle string
}

func NewDynamicStruct(table *unitable.Table) *DynamicStruct {
	var fields []reflect.StructField

	for _, col := range table.Columns {
		field := reflect.StructField{}
		field.Name = strcase.ToCamel(col.Name)
		switch col.Type {
		case "integer":
			field.Type = reflect.TypeOf(int64(0))
		case "float":
			field.Type = reflect.TypeOf(float64(0))
		case "string":
			switch col.Format {
			case "datetime", "time":
				now := time.Now()
				field.Type = reflect.TypeOf(&now)
			case "geometry":
				geo := &geom.Geometry{}
				field.Type = reflect.TypeOf(geo)
			default:
				if col.Repeated {
					field.Type = reflect.TypeOf([]string{})
				} else {
					field.Type = reflect.TypeOf("")
				}
			}
		}

		switch strings.ToLower(table.JsonStyle) {
		case LowerCamel, strings.ToLower(LowerCamel):
			field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, strcase.ToLowerCamel(col.Name)))
		default:
			field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, col.Name))
		}

		var gtags []string
		gtags = append(gtags, fmt.Sprintf("column:%s", col.Name))

		if col.Indexed {
			gtags = append(gtags, "index")
		}
		if len(col.DisplayName) > 0 {
			gtags = append(gtags, fmt.Sprintf("comment:%s", col.DisplayName))
		}
		if len(gtags) > 0 {
			tag := fmt.Sprintf(` gorm:"%s"`, strings.Join(gtags, ";"))
			field.Tag = reflect.StructTag(string(field.Tag) + tag)
		}

		fields = append(fields, field)
	}

	return &DynamicStruct{
		Fields:    fields,
		JsonStyle: table.JsonStyle,
	}
}

func NewDynamicStructWith(qry *query.Query, meta *MetaTable) *DynamicStruct {
	if len(qry.Fields) == 0 {
		return NewDynamicStruct(meta.Table)
	}

	columnIndex := meta.Table.ColumnIndex()
	table := &unitable.Table{}
	for name, field := range qry.Fields {
		c := &unitable.Column{
			Name:         name,
			OriginalName: field.GetName(),
		}

		if len(field.Projection.Functions) == 0 {
			if ci, ok := columnIndex[c.OriginalName]; ok {
				c.Type = ci.Type
				c.Format = ci.Format
			}
		} else {
			fun := field.GetFunction()
			switch fun {
			case "count":
				c.Type = "integer"
			case "sum", "avg", "min", "max":
				if ci, ok := columnIndex[c.OriginalName]; ok {
					c.Type = ci.Type
					c.Format = ci.Format
				}
			default:
				c.Type = "string"
			}
		}
		table.Columns = append(table.Columns, c)
	}
	return NewDynamicStruct(table)
}

func (s *DynamicStruct) GetType() reflect.Type {
	if s.Type == nil {
		s.Type = reflect.StructOf(s.Fields)
	}
	return s.Type
}

func (s *DynamicStruct) New() interface{} {
	return reflect.New(s.GetType()).Interface()
}

func (s *DynamicStruct) NewOf(object *core.Object) (interface{}, error) {
	instance := s.New()

	var json []byte
	var err error

	if s.JsonStyle == LowerCamel {
		json, err = jsoniter.Marshal(object.ToLowerCamelKeys())
	} else {
		json, err = jsoniter.Marshal(object.ToSnakeKeys())
	}

	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(json, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func ParseObject(value interface{}) (*core.Object, error) {
	json, err := jsoniter.Marshal(value)
	if err != nil {
		return nil, err
	}

	obj := &core.Object{}
	err = jsoniter.Unmarshal(json, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
