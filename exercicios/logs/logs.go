package exercicios

import  "strings"

// Application identifies the application that emitted a given log line.
func Application(logLine string) string {
	recommendationChar := '❗' // U+2757
	searchChar := '🔍'          // U+1F50D
	weatherChar := '☀'         // U+2600

	for _, char := range logLine {
		switch char {
		case recommendationChar:
			return "recommendation"
		case searchChar:
			return "search"
		case weatherChar:
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of a corrupted character with the original value.
func Replace(logLine string, corruptedChar, originalValue rune) string {
	return strings.ReplaceAll(logLine, string(corruptedChar), string(originalValue))
}

// WithinLimit checks if a log line is within a specific character limit.
func WithinLimit(logLine string, limit int) bool {
	count := 0
	for range logLine {
		count++
	}

	return count <= limit
}