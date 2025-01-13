package arrayProblems

import (
	"reflect"
	"testing"
)

func TestMultipleLeftRotations(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3}, [][]int{{3, 4, 5, 1, 2}, {4, 5, 1, 2, 3}}},
	}

	for _, test := range tests {
		result := MultipleLeftRotations(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("MultipleLeftRotations(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
