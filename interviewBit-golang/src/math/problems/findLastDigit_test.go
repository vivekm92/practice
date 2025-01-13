package mathProblems

import (
	"testing"
)

func TestFindLastDigit(t *testing.T) {
	tests := []struct {
		A        string
		B        string
		Expected int
	}{
		{"2", "10", 4},
		{"2", "0", 1},
		{"11231231", "1231241241", 1},
	}

	for _, test := range tests {
		result := FindLastDigit(test.A, test.B)
		if result != test.Expected {
			t.Errorf("FindLastDigit(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
