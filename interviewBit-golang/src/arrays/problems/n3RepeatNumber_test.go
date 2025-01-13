package arrayProblems

import "testing"

func TestRepeatNumber(t *testing.T) {

	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 1, 2, 3}, 1},
		{[]int{1, 2, 3, 1, 1}, 1},
	}

	for _, test := range tests {
		result := RepeatNumber(test.input)
		if result != test.output {
			t.Errorf("RepeatNumber(%v) = %v ; want %v", test.input, result, test.output)
		}
	}

}
