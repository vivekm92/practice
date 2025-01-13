package backtracking

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {

	tests := []struct {
		A        []int
		Expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}

	for _, test := range tests {
		result := Permutations(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Permutations(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
