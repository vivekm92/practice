package mathProblems

import (
	"testing"
)

func TestLcm(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected int64
	}{
		{4, 12, 12},
	}

	for _, test := range tests {
		result := Lcm(test.A, test.B)
		if result != test.Expected {
			t.Errorf("Lcm(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
