package twoPointers

import (
	"testing"
)

func TestAtMostTwoOccurrences(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 1, 1, 2, 2, 2}, 2},
	}

	for _, test := range tests {
		result := AtMostTwoOccurrences(test.A)
		if result != test.Expected {
			t.Errorf("AtMostTwoOccurrences(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
