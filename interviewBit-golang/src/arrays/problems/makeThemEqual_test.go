package arrayProblems

import "testing"

func TestMakeThemEqual(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{3, 1, 1, 3}, 2},
		{[]int{1, 1, 1}, 0},
		{[]int{6, 3, 1, 1, 3, 6}, 6},
	}

	for _, test := range tests {
		result := MakeThemEqual(test.A)
		if result != test.Expected {
			t.Errorf("MakeThemEqual(%v) = %v ; %v", test.A, result, test.Expected)
		}
	}
}
