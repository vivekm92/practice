package greedy

import (
	"testing"
)

func TestDistributeCandy(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 4, 5, 2, 1}, 13},
	}

	for _, test := range tests {
		result := DistributeCandy(test.A)
		if result != test.Expected {
			t.Errorf("DistributeCandy(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
