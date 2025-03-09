package palindrome

func IsPalindrome(input string) bool {
	runes := []rune(input)

	// Reverse slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed := string(runes)
	return reversed == input
}

func IsPalindromeTwoPointer(input string) bool {
	runes := []rune(input)
	left, right := 0, len(runes)-1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}
