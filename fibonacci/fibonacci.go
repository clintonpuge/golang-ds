package fibonacci

// Fibonacci recursion
func Fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return Fibonacci(num-1) + Fibonacci(num-2)
}
