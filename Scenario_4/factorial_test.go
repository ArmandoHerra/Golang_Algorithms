package factorial_test

import (
	"testing"

	"github.com/ArmandoHerra/Golang_Algorithms/factorial"
)

func TestFactorialRecursive(t *testing.T) {
	tests := []struct {
		input    int
		expected uint64
	}{
		{0, 1},
		{1, 1},
		{4, 24},
		{5, 120},
		{10, 3628800},
	}

	for _, tt := range tests {
		result := factorial.FactorialRecursive(tt.input)
		if result != tt.expected {
			t.Errorf("FactorialRecursive(%d) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
