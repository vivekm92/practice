package arrayExamples

import (
	"reflect"
	"testing"
)

func TestMaxSet(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 5, -7, 2, 3}, []int{1, 2, 5}},
		{[]int{-1, -1, -1, -1}, []int{}},
		{[]int{0, 0, -1, 0, 0}, []int{0, 0}},
		{[]int{0, 0, -1, 0, 0, 0}, []int{0, 0, 0}},
		{[]int{0, -1, 2, 3, -4, 100}, []int{100}},
	}

	for _, test := range tests {
		result := MaxSet(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("MaxSet(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}
