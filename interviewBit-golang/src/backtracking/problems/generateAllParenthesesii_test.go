package backtracking

import (
	"reflect"
	"testing"
)

func TestGenerateAllParenthesesii(t *testing.T) {
	tests := []struct {
		A        int
		Expected []string
	}{
		{0, []string{}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}

	for _, test := range tests {
		result := GenerateAllParenthesesii(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("GenerateAllParenthesesii(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
