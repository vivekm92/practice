package backtracking

import (
	"reflect"
	"testing"
)

func TestAllPossibleCombinations(t *testing.T) {
	tests := []struct {
		A        []string
		Expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"ab"}, []string{"a", "b"}},
		{[]string{"ab", "cd"}, []string{"ac", "ad", "bc", "bd"}},
	}

	for _, test := range tests {
		result := AllPossibleCombinations(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("AllPossibleCombinations(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
