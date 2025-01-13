package mathProblems

import (
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		A        int
		Expected string
	}{
		{1, "A"},
		{2, "B"},
		{26, "Z"},
		{28, "AB"},
		{52, "AZ"},
	}

	for _, test := range tests {
		result := ConvertToTitle(test.A)
		if result != test.Expected {
			t.Errorf("ConvertToTitle(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
