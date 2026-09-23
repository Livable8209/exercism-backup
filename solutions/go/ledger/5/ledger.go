package ledger

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var entriesCopy []Entry
	for _, entry := range entries {
		entriesCopy = append(entriesCopy, entry)
	}
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	// Just use the same sorting code from the previous iterations
	slices.SortFunc(entriesCopy, func(e1, e2 Entry) int {
		if compare := cmp.Compare(e1.Date, e2.Date); compare != 0 {
			return compare
		}
		if compare := cmp.Compare(e1.Description, e2.Description); compare != 0 {
			return compare
		}
		return cmp.Compare(e1.Change, e2.Change)
	})

	var finalLedger string // this could be a string builder actually

	formattedHeader, err := formatHeader(locale)
	if err != nil {
		log.Println(err.Error())
		return formattedHeader, err // formattedHeader is an empty string here
	}

	finalLedger += formattedHeader

	ss := make([]string, len(entriesCopy))
	for index, entry := range entriesCopy {
		string, err := formatRow(currency, locale, entry)
		if err != nil {
			return "", err
		}
		ss[index] = string
	}
	for i := range len(entriesCopy) {
		finalLedger += ss[i]
	}
	fmt.Println(finalLedger)
	return finalLedger, nil
}

func formatHeader(locale string) (string, error) {
	// Header formatting: also stolen from previous iterations, but now as one switch
	switch locale { // use switch statement cause gopls wont shut up about it
	case "nl-NL":
		return fmt.Sprintf("%-10s", "Datum") + " | " + fmt.Sprintf("%-25s", "Omschrijving") + " | " + fmt.Sprintf("%-13s", "Verandering") + "\n", nil
	case "en-US":
		return fmt.Sprintf("%-10s", "Date") + " | " + fmt.Sprintf("%-25s", "Description") + " | " + fmt.Sprintf("%-13s", "Change") + "\n", nil
	default:
		return "", errors.New("unknown locale when formatting header")
	}
}

func formatRow(currency string, locale string, entry Entry) (string, error) {
	// length check for dates, could be replaced with actual go date stuff, so automatic handling is there
	if len(entry.Date) != 10 {
		return "", errors.New("date is not exactly 10 characters long")
	}

	// This stuff can be replaced with whatever go's way of handling dates are
	year, firstDash, month, secondDash, day := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
	if firstDash != '-' {
		return "", errors.New("first dash is not present in date")
	}
	if secondDash != '-' {
		return "", errors.New("second dash is not present in date")
	}

	// Entry truncation
	entryDescription := entry.Description
	if len(entryDescription) > 25 {
		entryDescription = entryDescription[:22] + "..."
	} else {
		entryDescription = entryDescription + strings.Repeat(" ", 25-len(entryDescription))
	}

	var ledgerDate string
	if locale == "nl-NL" {
		ledgerDate = day + "-" + month + "-" + year
	} else if locale == "en-US" {
		ledgerDate = month + "/" + day + "/" + year
	}

	negative := false
	cents := entry.Change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	var ledgerTransaction string

	// a hell of alot of this code is duplicated, i should extract this into its own function then
	// only do checks for the locale specific changes.
	if locale == "nl-NL" {

		// this part should be its own function
		if currency == "EUR" {
			ledgerTransaction += "€"
		} else if currency == "USD" {
			ledgerTransaction += "$"
		} else {
			return "", errors.New("unknown curreny type in NL locale")
		}
		ledgerTransaction += " "

		// i think this can be done with some padding stuff, same sprintf stuff, just %0d2d instead of %-16s or whatever
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]

		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		if negative {
			ledgerTransaction += "-"
		}
		for i := len(parts) - 1; i >= 0; i-- {
			ledgerTransaction += parts[i] + "."
		}
		ledgerTransaction = ledgerTransaction[:len(ledgerTransaction)-1]
		ledgerTransaction += ","
		ledgerTransaction += centsStr[len(centsStr)-2:]
		ledgerTransaction += " "

	} else if locale == "en-US" {
		if negative {
			ledgerTransaction += "("
		}

		// this part also should be in its own function, like the one in the NL locale
		if currency == "EUR" {
			ledgerTransaction += "€"
		} else if currency == "USD" {
			ledgerTransaction += "$"
		} else {
			return "", errors.New("unknown currency type in US locale")
		}
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			ledgerTransaction += parts[i] + ","
		}
		ledgerTransaction = ledgerTransaction[:len(ledgerTransaction)-1]
		ledgerTransaction += "."
		ledgerTransaction += centsStr[len(centsStr)-2:]
		if negative {
			ledgerTransaction += ")"
		} else {
			ledgerTransaction += " "
		}
	} else {
		return "", errors.New("unknown locale when attempting to format transactions rows")
	}
	var al int // replace this bit with utf8 counting
	for range ledgerTransaction {
		al++
	}

	return ledgerDate + strings.Repeat(" ", 10-len(ledgerDate)) + " | " + entryDescription + " | " + strings.Repeat(" ", 13-al) + ledgerTransaction + "\n", nil

}
