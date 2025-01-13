package arrayProblems

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {

	tests := []struct {
		input  int
		output [][]int
	}{
		{0, [][]int{}},
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for _, test := range tests {
		result := PascalTriangle(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("PascalTriangle(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
