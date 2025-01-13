package arrayProblems

import (
	"testing"
)

func TestBribingWasNeverEasy(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{[]int{3, 1, 4, 2, 7, 8, 5, 6, 11, 10, 9}, []int{1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8}, 10},
		{[]int{2, 1, 5, 3, 4}, []int{2, 2, 2, 2, 2}, 3},
		{[]int{2, 5, 1, 3, 4}, []int{1, 2, 3, 2, 1}, -1},
	}

	for _, test := range tests {
		result := BribingWasNeverEasy(test.A, test.B)
		if result != test.Expected {
			t.Errorf("BribingWasNeverEasy(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
