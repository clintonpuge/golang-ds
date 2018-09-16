package mergesort

// MergeSort function
func MergeSort(numbers []int) {
	helper := make([]int, len(numbers))
	mergeSort(numbers, helper, 0, len(numbers)-1)
}

func mergeSort(numbers []int, helper []int, low int, high int) {
	if low < high {
		mid := int((low + high) / 2)
		mergeSort(numbers, helper, low, mid)
		mergeSort(numbers, helper, mid+1, high)
		merge(numbers, helper, low, mid, high)
	}
}

func merge(numbers []int, helper []int, low int, mid int, high int) {
	for i := low; i <= high; i++ {
		helper[i] = numbers[i]
	}
	leftI := low
	rightI := mid + 1
	current := low
	for leftI <= mid && rightI <= high {
		if helper[leftI] <= helper[rightI] {
			numbers[current] = helper[leftI]
			leftI++
		} else {
			numbers[current] = helper[rightI]
			rightI++
		}
		current++
	}
	remaining := mid - leftI
	for i := 0; i <= remaining; i++ {
		numbers[current+i] = helper[leftI+i]
	}
}
