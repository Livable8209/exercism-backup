package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) < 2 {
		return false
	}

	// from the solutions, turns out the ascii value can just be checked
	// so using an ascii table will help here

	totalValue := 0
	doubleIt := false

	for idx := len(id) - 1; idx >= 0; idx-- { // we go backwards so this makes sense
		if id[idx] < '0' || id[idx] > '9' { // single quotes actually have a different meaning? took me 2 days to figure that out
			return false
		}

		number := (int(id[idx] - '0')) // apparently this is the best way to take a byte and make it an int
		if doubleIt {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}

		totalValue += number
		doubleIt = !doubleIt
	}

	return totalValue%10 == 0
}
