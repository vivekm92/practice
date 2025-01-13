package mathProblems

import (
	"testing"
)

func TestLeapYeat(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{2020, 1},
		{700, 0},
		{800, 1},
		{2023, 0},
	}

	for _, test := range tests {
		result := LeapYear(test.A)
		if result != test.Expected {
			t.Errorf("LeapYear(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
