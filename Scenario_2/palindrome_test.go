package palindrome_test

import (
	"testing"

	"github.com/ArmandoHerra/Golang_Algorithms/palindrome"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"A", true},
		{"racecar", true},
		{"abba", true},
		{"hello", false},
		{"Madam", false}, // Case-sensitive scenario
	}

	for _, tt := range tests {
		result := palindrome.IsPalindrome(tt.input)
		if result != tt.expected {
			t.Errorf("isPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

func TestIsPalindromeTwoPointer(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"A", true},
		{"racecar", true},
		{"abba", true},
		{"hello", false},
		{"Madam", false}, // Case-sensitive scenario
	}

	for _, tt := range tests {
		result := palindrome.IsPalindromeTwoPointer(tt.input)
		if result != tt.expected {
			t.Errorf("IsPalindromeTwoPointer(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
