package mathProblems

import (
	"testing"
)

func TestPythagoreanTriplets(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{5, 1},
		{13, 3},
		{15, 4},
	}

	for _, test := range tests {
		result := PythagoreanTriplets(test.A)
		if result != test.Expected {
			t.Errorf("PythagoreanTriplets(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
