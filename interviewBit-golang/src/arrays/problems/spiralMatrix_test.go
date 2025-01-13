package arrayProblems

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		C        int
		Expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 5, [][]int{{1, 2, 3, 4, 5}}},
		{[]int{1, 2, 3, 4, 5}, 5, 1, [][]int{{1}, {2}, {3}, {4}, {5}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}

	for _, test := range tests {
		result := SpiralMatrix(test.A, test.B, test.C)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("SpiralMatrix(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
