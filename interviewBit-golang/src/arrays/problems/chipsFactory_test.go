package arrayProblems

import (
	"reflect"
	"testing"
)

func TestMoveZerosAtTheBack(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{0, 1, 2, 3}, []int{1, 2, 3, 0}},
		{[]int{0, 1, 2, 0}, []int{1, 2, 0, 0}},
	}

	for _, test := range tests {
		result := MoveZerosAtTheBack(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("MoveZerosAtTheBack(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
