package mathPrimers

import (
	"reflect"
	"testing"
)

func TestSquareSum(t *testing.T) {

	tests := []struct {
		A        int
		Expected [][]int
	}{
		{0, [][]int{}},
		{1, [][]int{}},
		{2, [][]int{{1, 1}}},
		{10, [][]int{{1, 3}}},
	}

	for _, test := range tests {
		result := SquareSum(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("SquareSum(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
