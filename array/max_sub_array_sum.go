package array

// MaxSubArraySum func
func MaxSubArraySum(array []int) int {
	maxSoFar := array[0]
	maxEndingHere := 0
	size := len(array)
	// -2, -1, -3, -4
	for i := 0; i < size; i++ {
		maxEndingHere += array[i]
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
