package techpalace

import (
	"log"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	finalstr := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	log.Println(finalstr)
	return finalstr
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var builder strings.Builder
	for range numStarsPerLine {
		builder.WriteString("*")
	}
	builder.WriteString("\n")
	builder.WriteString(welcomeMsg)
	builder.WriteString("\n")
	for range numStarsPerLine {
		builder.WriteString("*")
	}

	finalstr := builder.String()
	log.Println(finalstr)
	return finalstr
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	finalstr := strings.ReplaceAll(oldMsg, "*", "")
	finalstr = strings.TrimSpace(finalstr)

	return finalstr
}
