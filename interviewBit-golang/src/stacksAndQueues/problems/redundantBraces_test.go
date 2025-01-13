package stacksAndQueues

import (
	"testing"
)

func TestBraces(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"((a+b))", 1},
		{"(a+(b))", 1},
		{"((a*b)+(c+d))", 0},
	}

	for _, test := range tests {
		result := Braces(test.A)
		if result != test.Expected {
			t.Errorf("Braces(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
