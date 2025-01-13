package backtracking

import "testing"

func TestMaximalString(t *testing.T) {
	tests := []struct {
		A        string
		B        int
		Expected string
	}{
		{"524", 1, "542"},
		{"425", 1, "524"},
		{"425", 2, "542"},
	}

	for _, test := range tests {
		result := MaximalString(test.A, test.B)
		if result != test.Expected {
			t.Errorf("MaximalString(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
