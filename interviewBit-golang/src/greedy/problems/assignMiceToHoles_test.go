package greedy

import (
	"testing"
)

func TestAssignMiceToHoles(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{[]int{}, []int{}, 0},
		{[]int{-4, 0, 4}, []int{-2, 0, 4}, 2},
	}

	for _, test := range tests {
		result := AssignMiceToHoles(test.A, test.B)
		if result != test.Expected {
			t.Errorf("AssignMiceToHoles(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
