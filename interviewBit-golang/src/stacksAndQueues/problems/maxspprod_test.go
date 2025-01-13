package stacksAndQueues

import (
	"testing"
)

func TestMaxspprod(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 4, 3, 4}, 3},
		{[]int{10, 7, 100}, 0},
		{[]int{5, 9, 6, 8, 6, 4, 6, 9, 5, 4, 9}, 80},
	}

	for _, test := range tests {
		result := Maxspprod(test.A)
		if result != test.Expected {
			t.Errorf("Maxspprod(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
