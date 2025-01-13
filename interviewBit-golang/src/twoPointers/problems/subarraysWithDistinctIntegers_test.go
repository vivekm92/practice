package twoPointers

import "testing"

func TestSubarraysWithDistinctIntegers(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{1, 2, 1, 2, 3}, 2, 7},
	}

	for _, test := range tests {
		result := SubarraysWithDistinctIntegers(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SubarraysWithDistinctIntegers(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
