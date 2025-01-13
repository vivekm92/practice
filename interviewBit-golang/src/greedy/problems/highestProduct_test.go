package greedy

import (
	"testing"
)

func TestHighestProduct(t *testing.T) {

	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{2, 3, 4, -6, -5}, 120},
		{[]int{2, 3, 3, 6, 5}, 90},
	}

	for _, test := range tests {
		result := HighestProduct(test.A)
		if result != test.Expected {
			t.Errorf("HighestProduct(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
