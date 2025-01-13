package backtracking

import (
	"reflect"
	"testing"
)

func TestAllUniquePermutations(t *testing.T) {
	tests := []struct {
		A        []int
		Expected [][]int
	}{
		{[]int{1, 1, 2, 2}, [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
	}

	for _, test := range tests {
		result := AllUniquePermutations(test.A)
		for _, v1 := range result {
			flag := false
			for _, v2 := range test.Expected {
				if reflect.DeepEqual(v1, v2) {
					flag = true
					break
				}
			}
			if !flag {
				t.Errorf("AllUniquePermutations(%v) = %v ; want %v", test.A, result, test.Expected)
			}
		}

	}

}
