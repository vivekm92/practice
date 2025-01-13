package greedy

import (
	"testing"
)

func TestGasStation(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{[]int{1, 2}, []int{2, 1}, 1},
		{[]int{1, 2, 1}, []int{2, 1, 2}, -1},
	}

	for _, test := range tests {
		result := GasStation(test.A, test.B)
		if result != test.Expected {
			t.Errorf("GasStation(%v, %v) = %v : want %v", test.A, test.B, result, test.Expected)
		}
	}
}
