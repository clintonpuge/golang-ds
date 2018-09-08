package interpolation

// Search func
func Search(items []int, value int) int {
	start, end := 0, len(items)-1
	min, max := items[start], items[len(items)-1]
	for start <= end && value >= min && value <= max {
		pos := start + int((end-start)/(max-min))*(value-min)
		if value > items[pos] {
			start = pos + 1
		} else if value < items[pos] {
			end = pos - 1
		} else {
			return pos
		}
	}
	return -1
}
