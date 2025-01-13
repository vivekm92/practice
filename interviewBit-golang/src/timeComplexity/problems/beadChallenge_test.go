package timeComplexity

import (
	"testing"
)

func TestBeadChallenge(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{2, 3, 1}, 0},
		{[]int{100, 2, 50, 10, 1}, 1},
		{[]int{3, 3, 3}, 1},
	}

	for _, test := range tests {
		result := BeadChallenge(test.A)
		if result != test.Expected {
			t.Errorf("BeadChallenge(%v) = %v ; %v", test.A, result, test.Expected)
		}
	}
}
