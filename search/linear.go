package Search

// LinearSearch function
func LinearSearch(data []int, value int) bool {
	for _, item := range data {
		if item == value {
			return true
		}
	}
	return false
}
