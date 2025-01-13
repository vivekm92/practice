package greedy

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {

	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 1, 2}, 1},
		{[]int{2, 1, 2}, 2},
	}

	for _, test := range tests {
		result := MajorityElement(test.A)
		if result != test.Expected {
			t.Errorf("MajorityElement(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
