package search

import (
	"testing"
)

// TestBinarySearchGo test
func TestBinarySearchGo(t *testing.T) {
	tests := []struct {
		data   []int
		value  int
		result bool
	}{
		{data: []int{1, 3, 5, 7, 9}, value: 3, result: true},
		{data: []int{1, 3, 5, 7, 9}, value: 4, result: false},
	}

	for _, test := range tests {
		actual := BinarySearch(test.data, test.value)
		if actual != test.result {
			t.Errorf("It should return %t instead got %t", test.result, actual)
		}
	}
}
