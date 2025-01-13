package mathProblems

import (
	"reflect"
	"testing"
)

func TestKThPermutation(t *testing.T) {

	tests := []struct {
		A        int
		B        int64
		Expected []int
	}{
		{7, 7, []int{1, 2, 3, 5, 4, 6, 7}},
		{4, 17, []int{3, 4, 1, 2}},
		{3, 3, []int{2, 1, 3}},
	}

	for _, test := range tests {
		result := KThPermutation(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("KThPermutation(%v, %v) = %v ;	want %v", test.A, test.B, result, test.Expected)
		}
	}
}
