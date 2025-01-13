package mathProblems

import (
	"testing"
)

func TestNumberOfSundays(t *testing.T) {
	tests := []struct {
		A        string
		B        int
		Expected int
	}{
		{"Sunday", 1, 1},
		{"Friday", 1, 0},
		{"Friday", 1, 0},
		{"Thrusday", 10, 1},
	}

	for _, test := range tests {
		result := NumberOfSundays(test.A, test.B)
		if result != test.Expected {
			t.Errorf("NumberOfSundays(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
