package exercicios

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// Create a string with the desired number of stars
	starLine := strings.Repeat("*", numStarsPerLine)

	// Construct the final message with stars above and below the welcome message
	finalMessage := starLine + "\n" + welcomeMsg + "\n" + starLine

	return finalMessage
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// Remove all stars from the text
	noStarsMsg := strings.ReplaceAll(oldMsg, "*", "")

	// Remove leading and trailing whitespaces
	cleanedMsg := strings.TrimSpace(noStarsMsg)

	return cleanedMsg
}
