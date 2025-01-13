package stringProblems

import "testing"

func TestOowerOf2(t *testing.T) {

	tests := []struct {
		A        string
		Expected int
	}{
		{"128", 1},
		{"1", 0},
		{"0", 0},
		{"147573952589676412928", 1},
		{"12", 0},
	}

	for _, test := range tests {
		result := PowerOf2(test.A)
		if result != test.Expected {
			t.Errorf("PowerOf2(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
