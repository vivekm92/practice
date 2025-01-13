package backtracking

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected [][]int
	}{
		{[]int{2, 3}, 2, [][]int{{2}}},
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 4, 6, 8}, 8, [][]int{{2, 2, 2, 2}, {2, 2, 4}, {2, 6}, {4, 4}, {8}}},
		// {[]int{8, 10, 6, 11, 1, 16, 8}, 28, [][]int{{2, 2, 3}, {7}}},
	}

	for _, test := range tests {
		result := CombinationSum(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("CombinationSum(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
