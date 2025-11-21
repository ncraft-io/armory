package handlers

import "fmt"

func getColIndex(i int) string {
	a := 'A'
	if i < 26 {
		return fmt.Sprintf("%c", rune(int(a)+i))
	} else if i < 26*26 {
		d := i / 26
		rem := i % 26
		return fmt.Sprintf("%c%c", rune(int(a)+d-1), rune(int(a)+rem))
	} else {
		return ""
	}
}
