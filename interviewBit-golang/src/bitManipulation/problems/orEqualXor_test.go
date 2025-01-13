package bitManipulation

import "testing"

func TestOrEqualXor(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{0, 1}, {1, 3}, {4, 81}, {5, 243},
	}

	for _, test := range tests {
		result := OrEqualXor(test.A)
		if result != test.Expected {
			t.Errorf("OrEqualXor(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
