package factorial

func FactorialRecursive(num int) uint64 {
	if num <= 1 {
		return 1
	}
	return uint64(num) * FactorialRecursive(num-1)
}
