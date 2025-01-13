package arrayProblems

import (
	"testing"
)

func TestRepeatedNumber(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{3, 4, 1, 4, 2}, 4},
		{[]int{1, 2, 3}, -1},
		{[]int{3, 4, 1, 4, 1}, 4},
	}

	for _, test := range tests {
		result := RepeatedNumber(test.A)
		if result != test.Expected {
			t.Errorf("RepeatedNumber(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := RepeatedNumber1(test.A)
		if result != test.Expected {
			t.Errorf("RepeatedNumber1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
