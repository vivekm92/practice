package greedy

import "testing"

func TestBulbs(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{0, 1, 0, 1}, 4},
	}

	for _, test := range tests {
		result := Bulbs(test.A)
		if result != test.Expected {
			t.Errorf("Bulbs(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
