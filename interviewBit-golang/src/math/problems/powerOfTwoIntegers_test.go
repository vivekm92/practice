package mathProblems

import (
	"testing"
)

func TestIsPower(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{2, 0}, {4, 1}, {142, 0}, {1, 1}, {1024000000, 1},
	}

	for _, test := range tests {
		result := IsPower(test.A)
		if result != test.Expected {
			t.Errorf("IsPower(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
