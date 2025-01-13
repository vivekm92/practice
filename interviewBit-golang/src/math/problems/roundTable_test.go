package mathProblems

import (
	"testing"
)

func TestRoundTable(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{1, 2},
		{2, 2},
		{4, 12},
	}

	for _, test := range tests {
		result := RoundTable(test.A)
		if result != test.Expected {
			t.Errorf("RoundTable(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
