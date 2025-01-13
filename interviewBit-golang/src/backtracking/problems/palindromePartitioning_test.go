package backtracking

import (
	"reflect"
	"testing"
)

func TestPalindromePartitioning(t *testing.T) {
	tests := []struct {
		A        string
		Expected [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"efe", [][]string{{"e", "f", "e"}, {"efe"}}},
	}

	for _, test := range tests {
		result := PalindromePartitioning(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("PalindromePartitioning(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
