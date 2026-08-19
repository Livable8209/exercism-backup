package luhn

import (
	"fmt"
	"strconv"
	"unicode"
)

// guard clause
// walk through the string backwards somehow, insert all ints into their own slice and ye
// ok so unicode package exists, thats good
func Valid(id string) bool {
	if len(id) < 2 {
		return false
	}

	intsReversed := make([]int, 0, len(id)) // waste a few slots but that should be fine i think
	for i := len(id) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(id[i])) {
			retrivedInt, err := strconv.Atoi(string(id[i]))
			if err != nil {
				panic("something went horribly wrong when attempting to Atoi")
			}
			intsReversed = append(intsReversed, retrivedInt)
		}
	}

	// the doubling step thing
	// so thats why its called modulo 10
	// also i gave up on doing performant code so hooray
	// actually completing stuff :D
	intsDoubled := make([]int, 0, len(intsReversed))
	for i := range intsReversed {
		if i == 0 {
			continue
		}
		if i == 1 {
			intsDoubled = append(intsDoubled, intsReversed[i])
		}

		if i%2 == 0 {
			intDoubledAndModulo := (intsReversed[i] * 2) % 10
			intsDoubled = append(intsDoubled, intDoubledAndModulo)
		} else {
			intsDoubled = append(intsDoubled, intsReversed[i])
		}
	}

	result := 0
	for _, v := range intsDoubled {
		result += v
	}

	isValid := result%10 == 0

	fmt.Printf("id %v validity: %v\n", id, isValid)

	return isValid
}
