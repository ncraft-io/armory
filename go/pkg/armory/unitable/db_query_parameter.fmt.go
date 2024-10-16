package unitable

import (
	"github.com/gertd/go-pluralize"
	"strings"
)

func (x *DbQuery_Parameter) Format() string {
	if x != nil {
		if x.IsArray {
			if x.Type == "string" {
				if pluralize.NewClient().IsPlural(x.Name) {
					return x.Name
				} else {
					return x.Name + "::array"
				}
			} else {
				if pluralize.NewClient().IsPlural(x.Name) {
					return x.Name + ":" + x.Type
				} else {
					return x.Name + ":" + x.Type + ":array"
				}
			}
		} else {
			if x.Type == "string" {
				return x.Name
			} else {
				return x.Name + ":" + x.Type
			}
		}
	}
	return ""
}

func (x *DbQuery_Parameter) ToString() string {
	return x.Format()
}

func (x *DbQuery_Parameter) Parse(value string) error {
	if x != nil && len(value) > 0 {
		segments := strings.Split(value, ":")
		if len(segments) > 0 {
			x.Name = segments[0]
		}
		if len(segments) > 1 && len(segments[1]) > 0 {
			x.Type = strings.ToLower(segments[1])
		} else {
			x.Type = "string"
		}

		if len(segments) > 2 && len(segments[2]) > 0 {
			switch strings.ToLower(segments[2]) {
			case "true", "array":
				x.IsArray = true
			}
		} else {
			if pluralize.NewClient().IsPlural(x.Name) {
				x.IsArray = true
			}
		}
	}
	return nil
}

func ParseDbQueryParameter(value string) (*DbQuery_Parameter, error) {
	var v *DbQuery_Parameter
	if err := v.Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
