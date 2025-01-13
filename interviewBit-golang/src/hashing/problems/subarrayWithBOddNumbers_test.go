package hashing

import (
	"testing"
)

func TestSubarrayWithBOddNumbers(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{}, 0, 0},
		{[]int{4, 3, 2, 3, 4}, 2, 4},
		{[]int{5, 6, 7, 8, 9}, 3, 1},
	}

	for _, test := range tests {
		res := SubarrayWithBOddNumbersBruteForce(test.A, test.B)
		if res != test.Expected {
			t.Errorf("SubarrayWithBOddNumbersBruteForce(%v, %v) = %v ; want %v", test.A, test.B, res, test.Expected)
		}
	}

	for _, test := range tests {
		res := SubarrayWithBOddNumbers(test.A, test.B)
		if res != test.Expected {
			t.Errorf("SubarrayWithBOddNumbers(%v, %v) = %v ; want %v", test.A, test.B, res, test.Expected)
		}
	}
}
