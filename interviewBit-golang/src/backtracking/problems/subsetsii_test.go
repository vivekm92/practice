package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsetsii(t *testing.T) {
	tests := []struct {
		A        []int
		Expected [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{}, [][]int{{}}},
	}

	for _, test := range tests {
		result := Subsetsii(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Subsetsii(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
