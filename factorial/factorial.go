package recursion

// Factorial func
func Factorial(num int) int {
	// termination
	if num <= 1 {
		return 1
	}
	// body
	return num * Factorial(num-1)
}
