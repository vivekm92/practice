package mathProblems

import (
	"testing"
)

func TestSortedPermutationByRank(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"acb", 2},
		{"a", 1},
		{"string", 598},
	}

	for _, test := range tests {
		result := SortedPermutationRank(test.A)
		if result != test.Expected {
			t.Errorf("SortedPermutationRank(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
