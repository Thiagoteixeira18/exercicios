package count

import "strings"

func Count(s, w string) int {
	if w == "" {
		return 0 // If the word to count is empty, return 0
	}

	words := strings.Fields(s) // Split the string into words

	count := 0
	for _, word := range words {
		if word == w {
			count++
		}
	}

	return count
}

func main() {
	s := "apple orange banana apple mango apple"
	w := "apple"
	result := Count(s, w)
	println("Count of", w, "in", s, ":", result)
}
