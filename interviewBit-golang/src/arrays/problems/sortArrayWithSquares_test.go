package arrayProblems

import (
	"reflect"
	"testing"
)

func TestSortArrayWithSquares(t *testing.T) {

	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{-6, -3, -1, 2, 4, 5}, []int{1, 4, 9, 16, 25, 36}}, // Test Case 1
		{[]int{-5, -4, -2, 0, 1}, []int{0, 1, 4, 16, 25}},        // Test Case 2
	}

	for _, test := range tests {
		result := SortArrayWithSquares(test.input)
		if !reflect.DeepEqual(test.output, result) {
			t.Errorf("SortArrayWithSquares(%v) = %v; want %v", test.input, result, test.output)
		}
	}

	for _, test := range tests {
		result := SortArrayWithSquares1(test.input)
		if !reflect.DeepEqual(test.output, result) {
			t.Errorf("SortArrayWithSquares1(%v) = %v; want %v", test.input, result, test.output)
		}
	}

	for _, test := range tests {
		result := SortArrayWithSquares2(test.input)
		if !reflect.DeepEqual(test.output, result) {
			t.Errorf("SortArrayWithSquares2(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}
