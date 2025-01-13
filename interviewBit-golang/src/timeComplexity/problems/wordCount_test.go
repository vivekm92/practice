package timeComplexity

import "testing"

func TestWordCount(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"bonjour", 1},
		{"hasta la vista", 3},
		{" bonjour", 1},
		{" bonjour ", 1},
		{"bonjour ", 1},
		{" hasta la vista", 3},
		{" hasta la vista ", 3},
		{" hasta la vista   ", 3},
		{" hasta la vista    a", 4},
	}

	for _, test := range tests {
		result := WordCount(test.A)
		if result != test.Expected {
			t.Errorf("WordCount(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
