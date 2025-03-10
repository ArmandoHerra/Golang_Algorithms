package fibonacci

func FibonacciIterative(n int) int {
	if n < 2 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}
	return curr
}
