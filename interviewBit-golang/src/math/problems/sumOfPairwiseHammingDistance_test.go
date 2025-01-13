package mathProblems

import (
	"testing"
)

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{2, 4, 6}, 8},
		{[]int{1}, 0},
	}

	for _, test := range tests {
		result := HammingDistance(test.A)
		if result != test.Expected {
			t.Errorf("HammingDistance(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
