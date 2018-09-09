package bubblesort

import (
	"reflect"
	"testing"
)

// TestBinarySearch test
func TestBubbleSort(t *testing.T) {
	tests := []struct {
		data   []int
		result []int
	}{
		{data: []int{8, 3, 5, 4, 7, 9}, result: []int{3, 4, 5, 7, 8, 9}},
		{data: []int{8, 2, 5, 4, 7, 9, 1, 3}, result: []int{1, 2, 3, 4, 5, 7, 8, 9}},
	}

	for _, test := range tests {
		BubbleSort(test.data)
		if !reflect.DeepEqual(test.data, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, test.data)
		}
	}
}
