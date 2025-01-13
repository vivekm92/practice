package bsearch

import (
	"reflect"
	"testing"
)

func TestPrecompute(t *testing.T) {

	tests := []struct {
		input  int
		output int
	}{
		{2, 2}, {4, 8}, {10, 100}, {6, 36}, {12, 1728},
	}

	prod = ComputeProductOfDivisors()

	for _, test := range tests {
		result := prod[test.input]
		if result != test.output {
			t.Errorf("arr[%v] = %v; want %v", test.input, result, test.output)
		}
	}
}

func TestSimpleQueries(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{[]int{1, 2, 4}, []int{1, 2, 3, 4, 5, 6}, []int{8, 8, 8, 2, 2, 1}},
		{[]int{1, 3}, []int{1}, []int{3}},
		// {[]int{39, 99, 70, 24, 49, 13, 86, 43, 88, 74, 45, 92, 72, 71, 90, 32, 19, 76, 84, 46, 63, 15, 87, 1, 39, 58, 17, 65, 99, 43, 83, 29, 64, 67, 100, 14, 17, 100, 81, 26, 45, 40, 95, 94, 86, 2, 89, 57, 52, 91, 45}, []int{1221, 360, 459, 651, 958, 584, 345, 181, 536, 116, 1310, 403, 669, 1044, 1281, 711, 222, 280, 1255, 257, 811, 409, 698, 74, 838}, []int{}},
	}

	for _, test := range tests {
		result := SimpleQueries(test.input1, test.input2)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("SimpleQueries(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
