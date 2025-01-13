package mathProblems

import (
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{99, 9},
		{100, 1},
	}

	for _, test := range tests {
		result := DigitalRoot(test.A)
		if result != test.Expected {
			t.Errorf("DigitalRoot(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
