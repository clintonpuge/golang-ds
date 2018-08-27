package array

import (
	"reflect"
	"testing"
)

// TestRotateArray test for rotate
func TestRotateArray(t *testing.T) {
	tests := []struct {
		data   []int
		num    int
		result []int
	}{
		{data: []int{1, 2, 5, 8}, num: 2, result: []int{5, 8, 1, 2}},
		{data: []int{12, 2, 4, 18, 3}, num: 3, result: []int{18, 3, 12, 2, 4}},
	}
	for _, test := range tests {
		Rotate(test.data, test.num)
		if !reflect.DeepEqual(test.data, test.result) {
			t.Errorf("It should return %v instead got %v", test.result, test.data)
		}
	}
}
