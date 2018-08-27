package array

import (
	"testing"

	"github.com/clintonpuge/go-ds/helpers"
)

// TestMaxSubArraySum test
func TestMaxSubArraySum(t *testing.T) {
	tests := []struct {
		data   []int
		result int
	}{
		{data: []int{1, -2, 3, 4, -4, 6, -14, 8, 2}, result: 10},
		{data: []int{-2, -3, -1, -4}, result: -1},
		{data: []int{-5, 6, 7, 1, 4, -8, 16}, result: 26},
	}
	t.Logf("Should receive %v", helpers.CheckMark)
	for _, test := range tests {
		actual := MaxSubArraySum(test.data)
		if actual != test.result {
			t.Errorf("%v Should return %d instead got %d", helpers.Xmark, test.result, actual)
		} else {
			t.Logf("Should receive %d %v", test.result, helpers.CheckMark)
		}
	}
}
