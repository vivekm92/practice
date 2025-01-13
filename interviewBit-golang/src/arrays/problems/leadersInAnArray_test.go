package arrayProblems

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindLeaders(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{3, 2, 10, 7, 5, 2}, []int{2, 5, 7, 10}},
	}

	for _, test := range tests {
		result := FindLeaders(test.A)
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("FindLeaders(%v) = %v; want %v", test.A, result, test.Expected)
		}

	}
}
