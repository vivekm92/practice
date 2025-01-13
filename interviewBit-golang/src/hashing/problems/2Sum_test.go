package hashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := []struct {
		input1 []int
		input2 int
		output []int
	}{
		{[]int{2, 2, 3, 3, 4, 1}, 5, []int{1, 3}},
		{[]int{2, 2, 3, 3, 4, 1}, 7, []int{3, 5}},
		{[]int{2, 2, 3, 3, 4, 1}, 15, []int{}},
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 2, 3, 3, 4, 1}, 7, []int{3, 5}},
	}

	for _, test := range tests {
		result := TwoSum(test.input1, test.input2)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("TwoSum(%v, %v) = %v; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
