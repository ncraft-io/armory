package handlers

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/db/go/pkg/mojo/db"
	"github.com/mojo-lang/mojo/go/pkg/mojo/parser/syntax"
	"reflect"
)

func ParseQuery(request interface{}) (*db.Query, error) {
	if request == nil {
		return nil, nil
	}

	valueOf := reflect.Indirect(reflect.ValueOf(request))
	typeOf := valueOf.Type()

	if typeOf.Kind() == reflect.Struct {
		query := &db.Query{}

		if field, ok := typeOf.FieldByName("Filter"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsZero() {
				filter := reflect.Indirect(f).String()
				expr, err := syntax.ParseExpression(filter)
				if err != nil {
					return nil, core.NewInvalidArgumentError("invalid filter value: %s", filter)
				}
				query.Filter = expr
			}
		}

		if field, ok := typeOf.FieldByName("Order"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsNil() {
				if order, ok := reflect.Indirect(f).Interface().(*core.Ordering); ok {
					query.Order = order
				}
			}
		}

		if field, ok := typeOf.FieldByName("Unique"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsZero() {
				query.Unique = reflect.Indirect(f).Bool()
			}
		}

		if field, ok := typeOf.FieldByName("FieldMask"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsNil() {
				if fm, ok := reflect.Indirect(f).Interface().(*core.FieldMask); ok {
					query.FieldMask = fm
				}
			}
		}

		if field, ok := typeOf.FieldByName("PageSize"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsZero() {
				query.PageSize = int32(reflect.Indirect(f).Int())
			}
		}

		if field, ok := typeOf.FieldByName("PageToken"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsZero() {
				query.PageToken = reflect.Indirect(f).String()
			}
		}

		if field, ok := typeOf.FieldByName("Skip"); ok {
			f := valueOf.FieldByIndex(field.Index)
			if !f.IsZero() {
				query.Skip = int32(reflect.Indirect(f).Int())
			}
		}

		return query, nil
	}

	return nil, nil
}
