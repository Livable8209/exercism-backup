package isbnverifier

import (
	"log"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		log.Printf("Invalid ISBN: ISBN \"%v\" is not 10 digits", isbn)
		return false
	}

	// chekc using a loop wit the checking
	// like how i did with luhn's
	// but also a special case for X too since it outside the digit range

	for i := range len(isbn) { // step through the first 9? characters

		if isbn[i] < '0' || isbn[i] > '9' {
			log.Printf("Invalid ISBN: ISBN \"%v\" has invalid digits", isbn)
			return false
		}
		// digit := int(isbn[i] - '0')
		// fmt.Println(digit)
	}
	log.Printf("ISBN \"%v\" is valid!", isbn)

	return true
}
