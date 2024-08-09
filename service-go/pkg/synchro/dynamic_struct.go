package synchro

import (
	"fmt"
	"github.com/iancoleman/strcase"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/go/pkg/armory/unitable"
	"reflect"
	"time"
)

type DynamicStruct struct {
	Fields []reflect.StructField
	Type   reflect.Type
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
			default:
				field.Type = reflect.TypeOf("")
			}
		}

		field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, strcase.ToLowerCamel(col.Name)))
		if len(col.DisplayName) > 0 {
			field.Tag = reflect.StructTag(string(field.Tag) + fmt.Sprintf(` gorm:"comment:%s"`, col.DisplayName))
		}

		fields = append(fields, field)
	}

	return &DynamicStruct{
		Fields: fields,
	}
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

	json, err := jsoniter.Marshal(object)
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
