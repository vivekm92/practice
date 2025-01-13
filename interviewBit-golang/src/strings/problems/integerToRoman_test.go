package stringProblems

import "testing"

func TestIntegerToRoman(t *testing.T) {
	tests := []struct {
		A        int
		Expected string
	}{
		{14, "XIV"},
		{16, "XVI"},
		{41, "XLI"},
		{50, "L"},
	}

	for _, test := range tests {
		result := IntegerToRoman(test.A)
		if result != test.Expected {
			t.Errorf("IntegerToRoman(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
