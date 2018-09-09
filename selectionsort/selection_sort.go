package selectionsort

// 8, 3, 5, 4, 7, 9
// SelectionSort func
func SelectionSort(numbers []int) {
	for index := range numbers {
		minValue := numbers[index]
		minIndex := index
		for j := index; j < len(numbers)-1; j++ {
			if minValue > numbers[j+1] {
				minValue = numbers[j+1]
				minIndex = j + 1
			}
		}
		numbers[index], numbers[minIndex] = numbers[minIndex], numbers[index]
	}
}
