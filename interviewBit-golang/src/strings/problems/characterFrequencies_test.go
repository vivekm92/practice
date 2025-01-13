package stringProblems

import (
	"reflect"
	"testing"
)

func TestCharacterFrequencies(t *testing.T) {
	tests := []struct {
		A        string
		Expected []int
	}{
		{"interviewbit", []int{3, 1, 2, 2, 1, 1, 1, 1}},
		{"scaler", []int{1, 1, 1, 1, 1, 1}},
		{"vivek", []int{2, 1, 1, 1}},
	}

	for _, test := range tests {
		result := CharacterFrequencies(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("CharacterFrequencies(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
