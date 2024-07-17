package unitable

func (x *Column) IsTypeValid() bool {
	if x != nil && len(x.Type) > 0 {
		switch x.Type {
		case "bool", "integer", "float", "string":
			return true
		}
	}
	return false
}
