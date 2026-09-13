package handlers

import (
	"github.com/mojo-lang/mojo/go/pkg/compiler/mojo/parser/syntax"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/mojo-lang/mojo/go/pkg/mojo/db/query"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
	"regexp"
	"strconv"
)

// ParseQuery extracts the common request fields without assuming pointer kinds.
func ParseQuery(request interface{}) (*query.Query, error) {
	value := reflect.Indirect(reflect.ValueOf(request))
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return nil, core.NewInvalidArgumentError("request must be a non-nil struct")
	}
	q := &query.Query{}
	for _, name := range []string{"Filter", "Order", "Unique", "FieldMask", "PageSize", "PageToken", "Skip"} {
		field := value.FieldByName(name)
		if !field.IsValid() {
			continue
		}
		if !field.CanInterface() {
			return nil, core.NewInvalidArgumentError("inaccessible query field %s", name)
		}
		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		}
		switch name {
		case "Order":
			order, ok := value.FieldByName(name).Interface().(*core.Ordering)
			if !ok {
				return nil, core.NewInvalidArgumentError("invalid order type")
			}
			q.Order = proto.Clone(order).(*core.Ordering)
			for _, order := range q.Order.Orders {
				if order == nil || !queryFieldName.MatchString(order.Field) {
					return nil, core.NewInvalidArgumentError("invalid order field")
				}
			}
		case "FieldMask":
			mask, ok := value.FieldByName(name).Interface().(*core.FieldMask)
			if !ok {
				return nil, core.NewInvalidArgumentError("invalid field mask type")
			}
			q.FieldMask = proto.Clone(mask).(*core.FieldMask)
			for _, path := range q.FieldMask.Paths {
				if !queryFieldName.MatchString(path) {
					return nil, core.NewInvalidArgumentError("invalid field mask path")
				}
			}
		case "Filter", "PageToken":
			if field.Kind() != reflect.String {
				return nil, core.NewInvalidArgumentError("invalid query field %s", name)
			}
			if name == "PageToken" {
				q.PageToken = field.String()
			} else if field.String() != "" {
				expr, err := syntax.ParseExpression(field.String())
				if err != nil {
					return nil, core.NewInvalidArgumentError("invalid filter: %s", err)
				}
				q.Filter = expr
			}
		case "Unique":
			if field.Kind() != reflect.Bool {
				return nil, core.NewInvalidArgumentError("invalid unique type")
			}
			q.Unique = field.Bool()
		case "PageSize", "Skip":
			if !field.CanInt() {
				return nil, core.NewInvalidArgumentError("invalid pagination type")
			}
			n := field.Int()
			if n < 0 || n > math.MaxInt32 {
				return nil, core.NewInvalidArgumentError("invalid pagination value")
			}
			if name == "PageSize" {
				q.PageSize = int32(n)
			} else {
				q.Skip = int32(n)
			}
		}
	}
	if q.PageToken != "" {
		page, err := strconv.ParseInt(q.PageToken, 10, 32)
		if err != nil || page < 0 || q.PageSize <= 0 || page*int64(q.PageSize)+int64(q.Skip) > math.MaxInt32 {
			return nil, core.NewInvalidArgumentError("invalid page token")
		}
	}
	return q, nil
}

var queryFieldName = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_.]*$`)

func nextPageToken(q *query.Query, total int) string {
	if q == nil || q.PageSize <= 0 {
		return ""
	}
	page, _ := strconv.ParseInt(q.PageToken, 10, 64)
	if (page+1)*int64(q.PageSize)+int64(q.Skip) < int64(total) {
		return strconv.FormatInt(page+1, 10)
	}
	return ""
}
