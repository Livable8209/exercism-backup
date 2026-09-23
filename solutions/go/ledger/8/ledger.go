package ledger

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil { // i STILL dont know what this does
			return "", err
		}
	}

	var entriesCopy []Entry
	for _, entry := range entries {
		entriesCopy = append(entriesCopy, entry)
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

	var finalLedger strings.Builder // gopls made this a string builder automatically wow

	formattedHeader, err := formatHeader(locale)
	if err != nil {
		log.Println(err.Error())
		return formattedHeader, err // formattedHeader is an empty string here
	}

	finalLedger.WriteString(formattedHeader)

	formattedRows := make([]string, len(entriesCopy))
	for index, entry := range entriesCopy {
		string, err := formatRow(currency, locale, entry)
		if err != nil {
			return "", err
		}
		formattedRows[index] = string
	}
	for i := range len(entriesCopy) {
		finalLedger.WriteString(formattedRows[i])
	}

	fmt.Println(finalLedger.String())
	return finalLedger.String(), nil
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

func formatRow(currency, locale string, entry Entry) (string, error) {
	// Entry truncation
	entryDescription := entry.Description
	if len(entryDescription) > 25 {
		entryDescription = entryDescription[:22] + "..."
	} else {
		entryDescription = fmt.Sprintf("%-25s", entryDescription)
	}

	ledgerDate, err := formateDate(locale, entry)
	if err != nil {
		return ledgerDate, err
	}

	ledgerTransaction, err := formateTransaction(locale, currency, entry)
	if err != nil {
		return ledgerTransaction, err
	}

	descriptionSpacing := utf8.RuneCountInString(ledgerTransaction)

	return fmt.Sprintf("%-10s", ledgerDate) + " | " + entryDescription + " | " + strings.Repeat(" ", 13-descriptionSpacing) + ledgerTransaction + "\n", nil

}

func formateDate(locale string, entry Entry) (string, error) {
	// length check for dates, could be replaced with actual go date stuff, so automatic handling is there
	if len(entry.Date) != 10 {
		return "", errors.New("date is not exactly 10 characters long")
	}

	// dateFormatted, err := time.Parse(time.DateOnly, entry.Date)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return "", err
	// }

	// fmt.Println(dateFormatted, entry.Date)

	// This stuff can be replaced with whatever go's way of handling dates are
	year, firstDash, month, secondDash, day := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
	if firstDash != '-' {
		return "", errors.New("first dash is not present in date")
	}
	if secondDash != '-' {
		return "", errors.New("second dash is not present in date")
	}

	switch locale {
	case "nl-NL":
		return day + "-" + month + "-" + year, nil
	case "en-US":
		return month + "/" + day + "/" + year, nil
	default:
		return "", errors.New("unknown locale processing date")
	}
}

func formateTransaction(locale, currency string, entry Entry) (string, error) {
	// a hell of alot of this code is duplicated, i should extract this into its own function then
	// only do checks for the locale specific changes.
	// this entire thing also builds a string, so string builder would be much better

	var ledgerTransaction string

	negative := false
	cents := entry.Change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}

	switch locale {
	case "nl-NL":
		// this part should be its own function
		switch currency {
		case "EUR":
			ledgerTransaction += "€"
		case "USD":
			ledgerTransaction += "$"
		default:
			return "", errors.New("unknown curreny type in NL locale")
		}

		ledgerTransaction += " "

		// i think this can be done with some padding stuff, same sprintf stuff, just %0d2d instead of %-16s or whatever
		centsStr := strconv.Itoa(cents)
		fmt.Println(centsStr)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		fmt.Println(centsStr)
		rest := centsStr[:len(centsStr)-2]
		fmt.Println(rest)

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
	case "en-US":
		if negative {
			ledgerTransaction += "("
		}

		// this part also should be in its own function, like the one in the NL locale
		switch currency {
		case "EUR":
			ledgerTransaction += "€"
		case "USD":
			ledgerTransaction += "$"
		default:
			return "", errors.New("unknown currency type in US locale")
		}

		centsStr := strconv.Itoa(cents)
		fmt.Println(centsStr)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		fmt.Println(centsStr)
		rest := centsStr[:len(centsStr)-2]
		fmt.Println(rest)
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
	default:
		return "", errors.New("unknown locale when attempting to format transactions rows")
	}

	return ledgerTransaction, nil
}
