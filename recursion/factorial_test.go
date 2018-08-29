package recursion

import (
	"testing"

	"github.com/clintonpuge/go-ds/helpers"
)

// TestFactorial factorial recursion
func TestFactorial(t *testing.T) {
	tests := []struct {
		num    int
		result int
	}{
		{num: 5, result: 120},
		{num: 6, result: 720},
	}
	for _, test := range tests {
		actual := Factorial(test.num)
		if actual != test.result {
			t.Errorf("%v Should return %d instead got %d", helpers.Xmark, test.result, actual)
		} else {
			t.Logf("Should receive %d %v", test.result, helpers.CheckMark)
		}
	}
}
