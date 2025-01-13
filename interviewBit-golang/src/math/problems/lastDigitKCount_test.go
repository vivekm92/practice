package mathProblems

import (
	"testing"
)

func TestLastDigitKCount(t *testing.T) {

	tests := []struct {
		A        int
		B        int
		C        int
		Expected int
	}{
		{11, 111, 1, 11},
		{1, 1, 2, 0},
	}

	for _, test := range tests {
		result := LastDigitKCount(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("LastDigitKCount(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
