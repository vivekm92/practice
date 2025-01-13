package backtrackingExamples

import (
	"testing"
)

func TestMod(t *testing.T) {
	tests := []struct {
		input1 int
		input2 int
		input3 int
		output int
	}{
		{2, 3, 3, 2},
		{0, 3, 3, 0},
		{1, 200, 3, 1},
		{2, 0, 3, 1},
		{-1, 5, 20, 19},
	}

	for _, test := range tests {
		result := Mod(test.input1, test.input2, test.input3)
		if result != test.output {
			t.Errorf("Mod(%v, %v, %v) = %v; want %v", test.input1, test.input2, test.input3, result, test.output)
		}
	}
}
