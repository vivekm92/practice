package stringProblems

import (
	"testing"
)

func TestCountPalidromicWords(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"the fastest racecar", 1},
		{"wow mom", 2},
	}

	for _, test := range tests {
		result := CountPalidromicWords(test.A)
		if result != test.Expected {
			t.Errorf("CountPalidromicWords(%v) = %v; want %v", test.A, result, test.Expected)
		}
	}
}
