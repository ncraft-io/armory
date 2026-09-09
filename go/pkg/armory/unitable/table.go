package unitable

func (x *Table) ColumnIndex() map[string]*Column {
	if x != nil && len(x.Columns) > 0 {
		index := make(map[string]*Column)
		for _, col := range x.Columns {
			index[col.Name] = col
		}
		return index
	}
	return nil
}

func (x *Table) GetGeometryColumn() *Column {
	if x != nil {
		for _, column := range x.Columns {
			if column.Format == "geometry" && column.Type == "string" {
				return column
			}
		}
	}
	return nil
}
