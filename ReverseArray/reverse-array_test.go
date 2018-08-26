package ReverseArray

import (
	"reflect"
	"testing"
)

// TestReverseArray test
func TestReverseArray(t *testing.T) {
	tests := []struct {
		data   []int
		start  int
		end    int
		result []int
	}{
		{data: []int{1, 3, 5, 7, 9}, start: 0, end: 4, result: []int{9, 7, 5, 3, 1}},
		{data: []int{20, 3, 5, 21, 9}, start: 0, end: 4, result: []int{9, 21, 5, 3, 20}},
		{data: []int{3, 4, 5, 7, 20, 23}, start: 0, end: 5, result: []int{23, 20, 7, 5, 4, 3}},
	}

	for _, test := range tests {
		ReverseArray(test.data, test.start, test.end)
		if !reflect.DeepEqual(test.data, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, test.data)
		}
	}
}
