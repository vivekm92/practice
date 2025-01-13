package greedy

import (
	"reflect"
	"testing"
)

func TestMaximumXorWithMinX(t *testing.T) {

	tests := []struct {
		A        int
		Expected []int
	}{
		{1, []int{1, 1}},
		{2, []int{1, 2}},
		{7, []int{1, 6}},
		{11, []int{4, 11}},
	}

	for _, test := range tests {
		result := MaximumXorWithMinX(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("MaximumXorWithMinX(%v) = %v ; %v", test.A, result, test.Expected)
		}
	}
}
