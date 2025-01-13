package mathProblems

import (
	"testing"
)

func TestArmstrongNumber(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{2, 1}, {371, 1}, {123, 0},
	}

	for _, test := range tests {
		result := ArmstrongNumber(test.A)
		if result != test.Expected {
			t.Errorf("ArmstrongNumber(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
