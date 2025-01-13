package bsearch

import (
	"reflect"
	"testing"
)

func TestFirstIndex(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected []int
	}{
		{[]int{9, 1}, []int{7, 10, 3}, []int{0, -1, 0}},
		{[]int{2, 3, 4}, []int{2, 3, 4}, []int{0, 1, 2}},
	}

	for _, test := range tests {
		result := FirstIndex(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("FirstIndex(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
