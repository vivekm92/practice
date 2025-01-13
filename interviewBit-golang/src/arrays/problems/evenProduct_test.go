package arrayProblems

import "testing"

func TestEvenProduct(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 7},
		{[]int{1, 2, 3, 4}, 15},
		{[]int{1, 2, 3, 4, 5}, 31},
	}

	for _, test := range tests {
		result := EvenProduct(test.A)
		if result != test.Expected {
			t.Errorf("EvenProduct(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
