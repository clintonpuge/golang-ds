package mergesort

import (
	"reflect"
	"testing"
)

// TestMergeAndSort test
func TestMergeAndSort(t *testing.T) {
	tests := []struct {
		left   []int
		right  []int
		result []int
	}{
		{left: []int{8, 3, 5, 4, 7, 9}, right: []int{2, 1, 11, 6}, result: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}},
		{left: []int{8, 3, 5}, right: []int{2, 1, 11, 6}, result: []int{1, 2, 3, 5, 6, 8, 11}},
	}

	for _, test := range tests {
		MergeAndSort(test.data)
		if !reflect.DeepEqual(test.data, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, test.data)
		}
	}
}
