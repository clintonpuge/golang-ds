package Search

// BinarySearch fn
func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	var mid int

	for low <= high {
		mid = int((low + high) / 2)
		if value == data[mid] {
			return true
		}
		if value > data[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
