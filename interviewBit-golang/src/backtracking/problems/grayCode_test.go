package backtracking

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		A        int
		Expected []int
	}{
		{0, []int{}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 3, 2}},
		{3, []int{0, 1, 3, 2, 6, 7, 5, 4}},
		{4, []int{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}},
	}

	for _, test := range tests {
		result := GrayCode(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("GrayCode(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
