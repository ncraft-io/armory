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
	"sync"
	"time"
)

const LowerCamel = "lowerCamel"

type DynamicStruct struct {
	Fields    []reflect.StructField
	Type      reflect.Type
	typeOnce  sync.Once
	JsonStyle string
}

func NewDynamicStruct(table *unitable.Table) *DynamicStruct {
	var fields []reflect.StructField

	for _, col := range table.Columns {
		field := reflect.StructField{}
		field.Name = strcase.ToCamel(col.Name)
		switch col.Type {
		case "bool":
			field.Type = reflect.TypeOf(false)
		case "integer":
			field.Type = reflect.TypeOf(int64(0))
		case "float":
			field.Type = reflect.TypeOf(float64(0))
		case "string":
			switch col.Format {
			case "datetime", "time", "timestamp":
				now := time.Now()
				field.Type = reflect.TypeOf(&now)
			case "bytes":
				field.Type = reflect.TypeOf([]byte{})
			case "geometry":
				geo := &geom.Geometry{}
				field.Type = reflect.TypeOf(geo)
			default:
				if col.Repeated {
					field.Type = reflect.TypeOf(StringArray{})
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

		if col.Name == "id" {
			gtags = append(gtags, "primaryKey")
			if col.Type == "integer" {
				gtags = append(gtags, "autoIncrement:false")
			}
		}
		if col.Unique {
			gtags = append(gtags, "uniqueIndex")
		}
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
	if qry == nil || len(qry.Fields) == 0 {
		return NewDynamicStruct(meta.Table)
	}

	columnIndex := meta.Table.ColumnIndex()
	table := &unitable.Table{JsonStyle: meta.Table.JsonStyle}
	for name, field := range qry.Fields {
		c := &unitable.Column{
			Name:         name,
			OriginalName: field.GetName(),
		}

		if len(field.Projection.Functions) == 0 {
			if ci, ok := columnIndex[c.OriginalName]; ok {
				c.Type = ci.Type
				c.Format = ci.Format
				c.Repeated = ci.Repeated
			}
		} else {
			fun := field.GetFunction()
			switch fun {
			case "count":
				c.Type = "integer"
			case "avg":
				c.Type = "float"
			case "sum", "min", "max":
				if ci, ok := columnIndex[c.OriginalName]; ok {
					c.Type = ci.Type
					c.Format = ci.Format
					c.Repeated = ci.Repeated
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
	s.typeOnce.Do(func() {
		if s.Type == nil {
			s.Type = reflect.StructOf(s.Fields)
		}
	})
	return s.Type
}

func (s *DynamicStruct) New() interface{} {
	return reflect.New(s.GetType()).Interface()
}

func (s *DynamicStruct) NewSliceOf() interface{} {
	ptr := reflect.PointerTo(s.GetType())
	return reflect.New(reflect.SliceOf(ptr)).Interface()
}

func (s *DynamicStruct) NewOf(object *core.Object) (interface{}, error) {
	if object == nil {
		return nil, fmt.Errorf("nil row")
	}
	instance := s.New()

	var json []byte
	var err error

	if strings.EqualFold(s.JsonStyle, LowerCamel) {
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

func ParseObjects(value interface{}) ([]*core.Object, error) {
	json, err := jsoniter.Marshal(value)
	if err != nil {
		return nil, err
	}

	var objs []*core.Object
	err = jsoniter.Unmarshal(json, &objs)
	if err != nil {
		return nil, err
	}

	return objs, nil
}
