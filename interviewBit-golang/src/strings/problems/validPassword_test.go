package stringProblems

import "testing"

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"InterviewBit", 0},
		{"Scaler@1", 1},
	}

	for _, test := range tests {
		result := IsValidPassword(test.A)
		if result != test.Expected {
			t.Errorf("IsValidPassword(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
