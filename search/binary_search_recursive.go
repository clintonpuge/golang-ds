package search

// BinarySearchRecursive func
func BinarySearchRecursive(data []int, low int, high int, value int) int {
	if high < low {
		return -1
	}
	mid := int((low + high) / 2)
	if data[mid] > value {
		return BinarySearchRecursive(data, low, mid-1, value)
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return mid
	}
}
