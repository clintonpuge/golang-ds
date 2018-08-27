package array

// Rotate rotate array from left to right
func Rotate(data []int, num int) {
	Reverse(data, 0, num-1)
	Reverse(data, num, len(data)-1)
	Reverse(data, 0, len(data)-1)
}
