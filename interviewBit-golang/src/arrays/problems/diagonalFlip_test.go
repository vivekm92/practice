package arrayProblems

import (
	"reflect"
	"testing"
)

func TestDiagonalFlip(t *testing.T) {
	tests := []struct {
		A        [][]int
		Expected [][]int
	}{
		{[][]int{{1, 0}, {1, 0}}, [][]int{{1, 1}, {0, 0}}},
	}

	for _, test := range tests {
		result := DiagonalFlip(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("DiagonalFlip(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
