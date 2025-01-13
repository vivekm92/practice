package mathProblems

import "testing"

func TestSortedPermutationRankWithRepeats(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"aba", 2},
		{"bca", 4},
		{"bbbaaaaaa", 84},
		{"bbbbaaaaa", 126},
	}

	for _, test := range tests {
		result := SortedPermutationRankWithRepeats(test.A)
		if result != test.Expected {
			t.Errorf("SortedPermutationRankWithRepeats(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
