package mathProblems

import "testing"

func TestOddEvenRule(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		C        int
		Expected int
	}{
		{[]int{1, 2, 3}, 31, 100, 100},
		{[]int{0, 1, 1}, 2, 51, 102},
	}

	for _, test := range tests {
		result := OddEvenRule(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("OddEvenRule(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.C)
		}
	}
}
