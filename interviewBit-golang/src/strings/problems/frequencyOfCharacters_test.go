package stringProblems

import (
	"reflect"
	"testing"
)

func TestFrequencyOfCharacters(t *testing.T) {
	tests := []struct {
		A        string
		Expected []int
	}{
		{"abcdefghijklmnopqrstuvwxyz", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{"interviewbit", []int{0, 1, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 2, 0, 1, 1, 0, 0, 0}},
	}

	for _, test := range tests {
		result := FrequencyOfCharacters(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("FrequencyOfCharacters(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
