package gcd

import (
	"testing"

	"github.com/clintonpuge/go-ds/helpers"
)

// TestInterpolationSearch
func TestInterpolationSearch(t *testing.T) {
	tests := []struct {
		items  []int
		value  int
		result int
	}{
		{items: {1, 2, 9, 20, 31, 45, 63, 70, 100}, value: 63, result: 6},
	}
	for _, test := range tests {
		actual := InterpolationSearch(test.items, test.value)
		if actual != test.result {
			t.Errorf("%v Should return %d instead got %d", helpers.Xmark, test.result, actual)
		} else {
			t.Logf("Should receive %d %v", test.result, helpers.CheckMark)
		}
	}
}
