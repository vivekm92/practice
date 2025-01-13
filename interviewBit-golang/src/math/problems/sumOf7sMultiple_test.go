package mathProblems

import (
	"testing"
)

func TestSumOf7sMultiple(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected int64
	}{
		{1, 7, 7},
		{1, 1000000000, 71428571071428571},
	}

	for _, test := range tests {
		result := SumOf7sMultiple(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SumOf7sMultiple(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
