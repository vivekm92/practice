package backtracking

import (
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	tests := []struct {
		A        []int
		Expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
	}

	for _, test := range tests {
		result := Subset(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Subset(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
