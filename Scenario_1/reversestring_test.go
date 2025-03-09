package reversestring_test

import (
	"testing"

	"github.com/ArmandoHerra/Golang_Algorithms/reversestring"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"A", "A"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"12345", "54321"},
		{"Hello1", "1olleH"},
	}

	for _, tt := range tests {
		result := reversestring.ReverseStringRune(tt.input)
		if result != tt.expected {
			t.Errorf("ReverseStringRune(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}

	for _, tt := range tests {
		result2 := reversestring.ReverseStringBuilder(tt.input)
		if result2 != tt.expected {
			t.Errorf("ReverseStringBuilder(%q) = %q; want %q", tt.input, result2, tt.expected)
		}
	}
}
