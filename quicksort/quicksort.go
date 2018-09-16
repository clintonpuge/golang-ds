package quicksort

// QuickSort function
func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}
func quickSort(array []int, start int, end int) {
	if start <= end {
		pivotI := partition(array, start, end)
		quickSort(array, start, pivotI-1)
		quickSort(array, pivotI+1, end)
	}
}

func partition(array []int, start int, end int) int {
	pivot := array[end]
	pIndex := start
	for i := start; i <= end-1; i++ {
		if array[i] <= pivot {
			array[pIndex], array[i] = array[i], array[pIndex]
			pIndex++
		}
	}
	array[end], array[pIndex] = array[pIndex], array[end]
	return pIndex
}
