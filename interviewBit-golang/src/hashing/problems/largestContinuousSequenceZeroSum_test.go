package hashing

import (
	"reflect"
	"testing"
)

func TestLargestContinuousSequenceZeroSum(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{1, 2, -3, 3}, []int{1, 2, -3}},
		{[]int{1, 2, -2, 3, -3}, []int{2, -2, 3, -3}},
	}

	for _, test := range tests {
		result := LargestContinuousSequenceZeroSum(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("LargestContinuousSequenceZeroSum(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := LargestContinuousSequenceZeroSum1(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("LargestContinuousSequenceZeroSum1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
