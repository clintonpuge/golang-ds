package fibonacci

import (
	"testing"

	"github.com/clintonpuge/go-ds/helpers"
)

//
func TestGreatestCommonDivisor(t *testing.T) {
	tests := []struct {
		num    int
		result int
	}{
		{num: 6, result: 8},
		{num: 9, result: 34},
	}
	for _, test := range tests {
		actual := Fibonacci(test.num)
		if actual != test.result {
			t.Errorf("%v Should return %d instead got %d", helpers.Xmark, test.result, actual)
		} else {
			t.Logf("Should receive %d %v", test.result, helpers.CheckMark)
		}
	}
}
