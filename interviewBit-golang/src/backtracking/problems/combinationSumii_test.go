package backtracking

import (
	"reflect"
	"testing"
)

func TestCombinationSumii(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}

	for _, test := range tests {
		result := CombinationSumii(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("CombinationSumii(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
