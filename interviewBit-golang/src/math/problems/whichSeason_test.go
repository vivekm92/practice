package mathProblems

import "testing"

func TestWhichSeason(t *testing.T) {
	tests := []struct {
		A        int
		Expected string
	}{
		{2, "Winter"},
		{14, "Invalid"},
	}

	for _, test := range tests {
		result := WhichSeason(test.A)
		if result != test.Expected {
			t.Errorf("WhichSeason(%v) = %v ; %v", test.A, result, test.Expected)
		}
	}
}
