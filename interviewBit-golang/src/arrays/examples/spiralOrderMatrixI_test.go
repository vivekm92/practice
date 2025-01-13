package arrayExamples

import (
	"reflect"
	"testing"
)

func TestSpiralOrder1(t *testing.T) {

	tests := []struct {
		input  [][]int
		output []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}}, // Test Case 1
		{[][]int{{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},                             // Test Case 2
		{[][]int{{1}, {2}, {3}, {4}, {5}}, []int{1, 2, 3, 4, 5}},                     // Test Case 3
		{[][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 6, 5, 4}},                     // Test Case 4
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 4, 6, 5, 3}},                   // Test Case 5
	}

	for _, test := range tests {
		result := SpiralOrder1(test.input)
		if !reflect.DeepEqual(test.output, result) {
			t.Errorf("SpiralOrder(%v) = %v; want %v", test.input, result, test.output)
		}
	}

}
