package mathProblems

import (
	"testing"
)

func TestStepByStep(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{2, 3},
		{3, 2},
		{4, 3},
		{5, 5},
	}

	for _, test := range tests {
		result := StepByStep(test.A)
		if result != test.Expected {
			t.Errorf("StepByStep(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
