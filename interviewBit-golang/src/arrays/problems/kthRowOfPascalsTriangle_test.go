package arrayProblems

import (
	"reflect"
	"testing"
)

func TestKthRowOfPascalTriangle(t *testing.T) {
	tests := []struct {
		input  int
		output []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{13, []int{1, 13, 78, 286, 715, 1287, 1716, 1716, 1287, 715, 286, 78, 13, 1}},
	}

	for _, test := range tests {
		result := KthRowOfPascalTriangle(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("KthRowOfPascalTriangle(%v) = %v ; want %v", test.input, result, test.output)
		}
	}

	for _, test := range tests {
		result := kthRowOfPascalsTriangleRecursive(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("kthRowOfPascalsTriangleRecursive(%v) = %v ; want %v", test.input, result, test.output)
		}
	}

	for _, test := range tests {
		result := kthRowOfPascalsTriangleIterative(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("kthRowOfPascalsTriangleIterative(%v) = %v ; want %v", test.input, result, test.output)
		}
	}

	for _, test := range tests {
		result := kthRowOfPascalsTriangle1(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("kthRowOfPascalsTriangle1(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
