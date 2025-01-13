package stringProblems

import "testing"

func TestMultiplyStrings(t *testing.T) {
	tests := []struct {
		A        string
		B        string
		Expected string
	}{
		{"10", "12", "120"},
		{"11287654834672", "0", "0"},
	}

	for _, test := range tests {
		result := MultiplyStrings(test.A, test.B)
		if result != test.Expected {
			t.Errorf("MultiplyStrings(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
