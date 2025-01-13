package timeComplexity

import (
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		A        string
		Expected int64
	}{
		{"a12b34c", 46},
		{"123", 123},
	}

	for _, test := range tests {
		result := ExtractNumbers(test.A)
		if result != test.Expected {
			t.Errorf("ExtractNumbers(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
