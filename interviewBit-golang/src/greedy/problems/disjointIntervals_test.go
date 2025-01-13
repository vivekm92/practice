package greedy

import "testing"

func TestDisjointIntervals(t *testing.T) {

	tests := []struct {
		A       [][]int
		Expectd int
	}{
		{[][]int{{1, 4}, {1, 2}, {3, 4}, {5, 6}}, 3},
		{[][]int{{1, 3}, {2, 4}, {3, 5}}, 1},
	}

	for _, test := range tests {
		result := DisjointIntervals(test.A)
		if result != test.Expectd {
			t.Errorf("DisjointIntervals(%v) = %v ; want %v", test.A, result, test.Expectd)
		}
	}
}
