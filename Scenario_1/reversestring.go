package reversestring

import "strings"

func ReverseStringRune(input string) string {
	// Convert string into a slice of runes since strings in Golang are immutable.
	runes := []rune(input)

	// Reverse in-place
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert back into string
	return string(runes)
}

func ReverseStringBuilder(input string) string {
	var builder strings.Builder
	// Pre-allocate capacity for performance (optional)
	builder.Grow(len(input))

	// Append characters in reverse order
	for i := len(input) - 1; i >= 0; i-- {
		builder.WriteByte(input[i])
	}

	return builder.String()
}
