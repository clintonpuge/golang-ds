package mergesort

import (
	"reflect"
	"testing"
)

// TestMergeAndSort test
func TestMergeAndSort(t *testing.T) {
	left1 := make([]int, 3)
	left1 = append([]int{1, 2, 4}, left1...)
	tests := []struct {
		left       []int
		right      []int
		leftCount  int
		rightCount int
		result     []int
	}{
		{left: left1, right: []int{3, 5, 6}, result: []int{1, 2, 3, 4, 5, 6}, leftCount: 3, rightCount: 3},
	}

	for _, test := range tests {
		MergeAndSort(test.left, test.right, test.leftCount, test.rightCount)
		if !reflect.DeepEqual(test.left, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, test.left)
		}
	}
}
