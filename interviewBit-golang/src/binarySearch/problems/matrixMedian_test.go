package bsearch

import "testing"

func TestMatrixMedian(t *testing.T) {
	tests := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{1}, {3}, {3}}, 3},
		{[][]int{{1, 3, 5}, {2, 6, 9}, {3, 6, 9}}, 5},
		{[][]int{{5, 17, 100}}, 17},
	}

	for _, test := range tests {
		result := MatrixMedian(test.input)
		if result != test.output {
			t.Errorf("MatrixMedian(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}
