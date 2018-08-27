package array

// Reverse :
// step 1 swap start and end
// step 2 increment start and decrement end
func Reverse(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
