package gcd

import (
	"testing"

	"github.com/clintonpuge/go-ds/helpers"
)

//
func TestGreatestCommonDivisor(t *testing.T) {
	tests := []struct {
		num1   int
		num2   int
		result int
	}{
		{num1: 270, num2: 192, result: 6},
		{num1: 10, num2: 45, result: 5},
	}
	for _, test := range tests {
		actual := GreatestCommonDivisor(test.num1, test.num2)
		if actual != test.result {
			t.Errorf("%v Should return %d instead got %d", helpers.Xmark, test.result, actual)
		} else {
			t.Logf("Should receive %d %v", test.result, helpers.CheckMark)
		}
	}
}
