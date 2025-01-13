package hashing

import (
	"reflect"
	"testing"
)

func TestAnIncrementProblem(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{1, 1}, []int{2, 1}},
		{[]int{1, 2, 1, 2, 2}, []int{3, 3, 1, 2, 2}},
		{[]int{1, 1, 2, 2}, []int{3, 1, 3, 2}},
		{[]int{1, 2}, []int{1, 2}},
	}

	for _, test := range tests {
		result := AnIncrementProblem(test.A)
		if !reflect.DeepEqual(test.A, result) {
			t.Errorf("AnIncrementProblem(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	// for _, test := range tests {
	// 	result := AnIncrementProblem1(test.A)
	// 	if !reflect.DeepEqual(test.A, result) {
	// 		t.Errorf("AnIncrementProblem1(%v) = %v ; want %v", test.A, result, test.Expected)
	// 	}
	// }

}
