package mathProblems

import (
	"reflect"
	"testing"
)

func TestPrimeSum(t *testing.T) {
	tests := []struct {
		A        int
		Expected []int
	}{
		{28, []int{5, 23}},
		{20000001, []int{2, 19999999}},
	}

	for _, test := range tests {
		result := PrimeSum(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("PrimeSum(%v) = %v ; want %v", test.A, result, test.Expected)
		}

	}
}
