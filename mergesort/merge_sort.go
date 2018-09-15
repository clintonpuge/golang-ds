package mergesort

// MergeSort function
func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	mid := len(numbers) / 2
	leftArr := MergeSort(numbers[:mid])
	rightArr := MergeSort(numbers[mid:])
	return merge(leftArr, rightArr)
}

func merge(left []int, right []int) []int {
	arr := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(arr, right...)
		}
		if len(right) == 0 {
			return append(arr, left...)
		}
		if left[0] <= right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	return arr
}
