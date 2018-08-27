package search

// LinearSearch function
func Linear(data []int, value int) bool {
	for _, item := range data {
		if item == value {
			return true
		}
	}
	return false
}
