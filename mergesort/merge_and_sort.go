package mergesort

// MergeAndSort Merge two arrays and sort them
func MergeAndSort(left []int, right []int, leftLast int, rightLast int) {
	leftIndex := leftLast - 1
	rightIndex := rightLast - 1
	mergeI := leftLast + rightLast - 1
	for rightIndex >= 0 {
		if left[leftIndex] >= 0 && left[leftIndex] > right[rightIndex] {
			left[mergeI] = left[leftIndex]
			leftIndex--
		} else {
			left[mergeI] = right[rightIndex]
			rightIndex--
		}
		mergeI--
	}
}
