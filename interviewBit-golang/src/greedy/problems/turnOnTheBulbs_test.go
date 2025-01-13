package greedy

import (
	"testing"
)

func TestTurnOnTheBulbs(t *testing.T) {

	tests := []struct {
		A        int
		Expected int
	}{
		{1, 1},
		{2, 2},
		{85008, 28336},
	}

	for _, test := range tests {
		result := TurnOnTheBulbs(test.A)
		if result != test.Expected {
			t.Errorf("TurnOnTheBulbs(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
