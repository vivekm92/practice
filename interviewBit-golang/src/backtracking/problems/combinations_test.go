package backtracking

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, test := range tests {
		result := Combinations(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Combinations(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
