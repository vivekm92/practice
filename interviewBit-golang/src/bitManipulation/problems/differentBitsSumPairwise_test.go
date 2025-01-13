package bitManipulation

import "testing"

func TestDifferentBitsSumPairwise(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 3, 5}, 8},
		{[]int{2, 3}, 2},
	}
	for _, test := range tests {
		result := DifferentBitsSumPairwise(test.A)
		if result != test.Expected {
			t.Errorf("DifferentBitsSumPairwise(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
