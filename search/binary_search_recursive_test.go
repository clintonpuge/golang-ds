package search

import (
	"testing"
)

// TestBinarySearchRecursive test
func TestBinarySearchRecursive(t *testing.T) {
	tests := []struct {
		data   []int
		value  int
		result int
	}{
		{data: []int{1, 3, 5, 7, 9}, value: 7, result: 3},
		{data: []int{1, 3, 5, 7, 9}, value: 4, result: -1},
	}

	for _, test := range tests {
		actual := BinarySearchRecursive(test.data, 0, 4, test.value)
		if actual != test.result {
			t.Errorf("It should return %d instead got %d", test.result, actual)
		}
	}
}
