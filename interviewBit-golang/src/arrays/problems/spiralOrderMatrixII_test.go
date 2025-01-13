package arrayProblems

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {

	tests := []struct {
		input  int
		output [][]int
	}{
		{5, [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}},
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {4, 3}}},
	}

	for _, test := range tests {
		result := GenerateMatrix(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("GenerateMatrix(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
