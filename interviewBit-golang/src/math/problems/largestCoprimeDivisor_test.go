package mathProblems

import (
	"testing"
)

func TestLargestCoprimeDivisor(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected int
	}{
		{30, 12, 5},
	}

	for _, test := range tests {
		result := LargestCoprimeDivisor(test.A, test.B)
		if result != test.Expected {
			t.Errorf("LargestCoprimeDivisor(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
