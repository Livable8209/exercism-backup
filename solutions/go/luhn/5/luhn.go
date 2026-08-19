package luhn

import (
	"fmt"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) < 2 {
		return false
	}

	// from the solutions, turns out the ascii value can just be checked
	// so using an ascii table will help here

	// totalValue := 0

	for idx := range 10 {
		fmt.Printf("%d=%d\n", idx, int('0')+idx)
	}
	return false
}
