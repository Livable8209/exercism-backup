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
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil { // wtf does this even do?
			return "", err
		}
	}

	slices.SortFunc(entriesCopy, func(e1, e2 Entry) int {
		if compare := cmp.Compare(e1.Date, e2.Date); compare != 0 {
			return compare
		}
		if compare := cmp.Compare(e1.Description, e2.Description); compare != 0 {
			return compare
		}

		return cmp.Compare(e1.Change, e2.Change)
	})

	formattedLedger, err := formatLedgerWithLocale(locale)
	if err != nil {
		return formattedLedger, err
	}

	for _, entry := range entriesCopy {
		if len(entry.Date) != 10 {
			log.Println("Ledger is not exactly 10 characters:", entry.Date)
			return "", errors.New("Ledger date isn't exactly 10 characters")
		}

		switch locale {
		case "en-US":
			log.Println("en-US")
		case "nl-NL":
			log.Println("nl-NL")
		default:
			log.Println("Unknown locale")
			return "", errors.New("Unknown locale")
		}

	}

	ss := make([]string, len(entriesCopy))
	for i, entry := range entriesCopy {
		vi, vs, ve := bigUglyFunc(i, currency, locale, entry)
		if ve != nil {
			return "", ve
		}
		ss[vi] = vs
	}
	for i := range len(entriesCopy) {
		formattedLedger += ss[i]
	}
	return formattedLedger, nil
}

func formatLedgerWithLocale(locale string) (string, error) {
	translatedLedger := ""
	switch locale { // use switch statement cause gopls wont shut up about it
	case "nl-NL":
		translatedLedger = fmt.Sprintf("%-10s", "Datum") +
			" | " +
			fmt.Sprintf("%-25s", "Omschrijving") +
			" | " + fmt.Sprintf("%-13s", "Verandering") + "\n"

	case "en-US":
		translatedLedger = fmt.Sprintf("%-10s", "Date") +
			" | " +
			fmt.Sprintf("%-25s", "Description") +
			" | " + fmt.Sprintf("%-13s", "Change") + "\n"

	}

	if translatedLedger != "" {
		return translatedLedger, nil
	} else {
		return translatedLedger, errors.New("Unable to format Ledger.")
	}
}

func bigUglyFunc(i int, currency, locale string, entry Entry) (int, string, error) {
	if len(entry.Date) != 10 {
		return i, "", errors.New("orig: bad entry length")
	}
	d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
	if d2 != '-' {
		return i, "", errors.New("orig: bad date seperator #1")
	}
	if d4 != '-' {
		return i, "", errors.New("orig: bad date seperator #2")
	}
	de := entry.Description
	if len(de) > 25 {
		de = de[:22] + "..."
	} else {
		de = de + strings.Repeat(" ", 25-len(de))
	}
	var d string
	if locale == "nl-NL" {
		d = d5 + "-" + d3 + "-" + d1
	} else if locale == "en-US" {
		d = d3 + "/" + d5 + "/" + d1
	}
	negative := false
	cents := entry.Change
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	var a string
	if locale == "nl-NL" {
		if currency == "EUR" {
			a += "€"
		} else if currency == "USD" {
			a += "$"
		} else {
			return i, "", errors.New("orig: bad currency type 1")
		}
		a += " "
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
			a += "-"
		}
		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + "."
		}
		a = a[:len(a)-1]
		a += ","
		a += centsStr[len(centsStr)-2:]
		a += " "
	} else if locale == "en-US" {
		if negative {
			a += "("
		}
		if currency == "EUR" {
			a += "€"
		} else if currency == "USD" {
			a += "$"
		} else {
			return i, "", errors.New("orig: bad currency type 2")
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
			a += parts[i] + ","
		}
		a = a[:len(a)-1]
		a += "."
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += ")"
		} else {
			a += " "
		}
	} else {
		return i, "", errors.New("orig: bad locale")
	}
	var al int
	for range a {
		al++

	}
	return i,
		d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
			strings.Repeat(" ", 13-al) + a + "\n", nil
}
