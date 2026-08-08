package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var finalString strings.Builder
	if number%3 == 0 {
		finalString.WriteString("Pling")
	}
	if number%5 == 0 {
		finalString.WriteString("Plang")
	}
	if number%7 == 0 {
		finalString.WriteString("Plong")
	}

	if finalString.Len() == 0 {
		finalString.WriteString(strconv.Itoa(number))
	}
	return finalString.String()
}
