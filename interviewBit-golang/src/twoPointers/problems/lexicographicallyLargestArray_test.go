package twoPointers

import (
	"reflect"
	"testing"
)

func TestLexicographicallyLargestArray(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{10, 20, 30, 40}, []int{40, 30, 20, 10}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 1}},
		{[]int{4, 1, 2, 3}, []int{4, 3, 2, 1}},
	}

	for _, test := range tests {
		result := LexicographicallyLargestArray(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("LexicographicallyLargestArray(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
