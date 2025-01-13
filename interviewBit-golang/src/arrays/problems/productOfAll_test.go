package arrayProblems

import (
	"reflect"
	"testing"
)

func TestProductOfAll(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{9, 9, 9}, []int{81, 81, 81}},
	}

	for _, test := range tests {
		result := ProductOfAll(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("ProductOfAll(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
