package reversestring

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
