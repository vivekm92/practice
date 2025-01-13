package timeComplexity

import (
	"testing"
)

func TestCheckIfPanagram(t *testing.T) {
	tests := []struct {
		A        []string
		Expected int
	}{
		{[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}, 1},
		{[]string{"bit", "scale"}, 0},
	}

	for _, test := range tests {
		result := CheckIfPanagram(test.A)
		if result != test.Expected {
			t.Errorf("CheckIfPanagram(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
