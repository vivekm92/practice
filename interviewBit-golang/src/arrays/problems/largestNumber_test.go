package arrayProblems

import (
	"reflect"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		input  []int
		output string
	}{
		{[]int{2, 10}, "210"},               // Test case 1
		{[]int{10, 2}, "210"},               // Test case 2
		{[]int{0, 0, 0, 0}, "0"},            // Test case 3
		{[]int{3, 30, 34, 5, 9}, "9534330"}, // Test case 4
	}

	for _, test := range tests {
		result := LargestNumber(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("LargestNumber(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}
