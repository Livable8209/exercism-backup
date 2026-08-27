package isbnverifier

import (
	"log"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbnStripped := strings.ReplaceAll(isbn, "-", "")

	if len(isbnStripped) != 10 {
		log.Printf("Invalid ISBN: ISBN \"%v\" is not 10 digits", isbn)
		return false
	}
	// chekc using a loop wit the checking
	// like how i did with luhn's
	// but also a special case for X too since it outside the digit range

	finalValue := 0
	for i := range len(isbnStripped) - 1 { // step through the first 9 characters because X isnt a digit
		if isbnStripped[i] < '0' || isbnStripped[i] > '9' {
			log.Printf("Invalid ISBN: ISBN \"%v\" has invalid digits", isbn)
			return false
		}
		digit := int(isbnStripped[i] - '0')
		finalValue += digit * (10 - i) // i think this works
		// turns out it DIDNT work because 11 is not 10, i hate off by one errors >:(
	}

	lastCharOrDigit := isbnStripped[len(isbnStripped)-1]

	if lastCharOrDigit == 'X' {
		finalValue += 10
	} else if !(lastCharOrDigit < '0' || lastCharOrDigit > '9') {
		finalValue += int(lastCharOrDigit - '0')
	} else {
		log.Printf("Invalid ISBN: ISBN \"%v\" has invalid check digit", isbn)
		return false
	}

	isValid := finalValue%11 == 0

	if isValid {
		log.Printf("ISBN \"%v\" is valid!", isbn)
	} else {
		log.Printf("Invalid ISBN: ISBN \"%v\" fails verification", isbn)
	}

	return isValid
}
