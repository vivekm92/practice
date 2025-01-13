package hashing

import (
	"testing"
)

func TestLongestSubarrayLength(t *testing.T) {

	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 0}, 3},
		{[]int{0, 1, 1, 0, 1}, 5},
	}

	for _, test := range tests {
		result := LongestSubarrayLength(test.A)
		if result != test.Expected {
			t.Errorf("LongestSubarrayLength(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := LongestSubarrayLengthBruteForce(test.A)
		if result != test.Expected {
			t.Errorf("LongestSubarrayLengthBruteForce(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
