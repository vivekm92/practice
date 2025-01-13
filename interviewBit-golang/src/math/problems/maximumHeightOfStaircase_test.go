package mathProblems

import (
	"testing"
)

func TestMaxStaircaseHeight(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 2}, {5, 2}, {6, 3}, {7, 3}, {8, 3}, {9, 3},
		{10, 4}, {11, 4}, {12, 4}, {13, 4}, {14, 4}, {15, 5}, {16, 5}, {17, 5}, {18, 5}, {19, 5},
	}

	for _, test := range tests {
		result := MaxStaircaseHeight(test.A)
		if result != test.Expected {
			t.Errorf("MaxStaircaseHeight(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
