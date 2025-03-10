package fibonacci_test

import (
	"testing"

	"github.com/ArmandoHerra/Golang_Algorithms/fibonacci"
)

func TestFibbonacciIterative(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{15, 610},
	}

	for _, tt := range tests {
		result := fibonacci.FibonacciIterative(tt.input)
		if result != tt.expected {
			t.Errorf("FibonacciIterative(%d) = %d; want %d", tt.input, result, tt.expected)
		}
	}
}
