package backtracking

import (
	"reflect"
	"testing"
)

func TestLetterPhone(t *testing.T) {
	tests := []struct {
		A        string
		Expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"0", []string{"0"}},
		{"7", []string{"p", "q", "r", "s"}},
	}

	for _, test := range tests {
		result := LetterPhone(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("LetterPhone(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
