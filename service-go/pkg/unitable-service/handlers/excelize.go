package handlers

import "github.com/xuri/excelize/v2"

func getColIndex(i int) string {
	name, _ := excelize.ColumnNumberToName(i + 1)
	return name
}
