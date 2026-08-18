package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

// guard clause
// strip spaes with fields
// ok use fieldsfunc instead because more flexible
func Valid(id string) bool {
	if len(id) < 2 {
		return false
	}

	idStripped := strings.FieldsFunc(id, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	fmt.Println(id, idStripped)

	return true
}
