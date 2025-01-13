package twoPointers

import (
	"reflect"
	"testing"
)

func TestMaxContinuos1s(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected []int
	}{
		{[]int{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}, 1, []int{0, 1, 2, 3, 4}},
		{[]int{1, 0, 0, 0, 1, 0, 1}, 2, []int{3, 4, 5, 6}},
	}

	for _, test := range tests {
		result := MaxContinuos1s(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("MaxContinuos1s(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
