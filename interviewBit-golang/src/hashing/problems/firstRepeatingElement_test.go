package hashing

import (
	"testing"
)

func TestFirstRepeatingElement(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 2, 2, 4, 4, 1}, 1},
	}

	for _, test := range tests {
		result := FirstRepeatingElement(test.A)
		if result != test.Expected {
			t.Errorf("FirstRepeatingElement(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := FirstRepeatingElement1(test.A)
		if result != test.Expected {
			t.Errorf("FirstRepeatingElement1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
