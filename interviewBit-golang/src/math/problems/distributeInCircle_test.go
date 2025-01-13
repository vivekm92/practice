package mathProblems

import (
	"testing"
)

func TestDistributeInCircle(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		C        int
		Expected int
	}{
		{2, 5, 1, 2},
		{8, 5, 2, 4},
	}

	for _, test := range tests {
		result := DistributeInCircle(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("DistributeInCircle(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
