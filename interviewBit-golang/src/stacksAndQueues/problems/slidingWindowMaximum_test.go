package stacksAndQueues

import (
	"reflect"
	"testing"
)

func TestSlidingMaximum(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 1, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		result := SlidingMaximum(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("SlidingMaximum(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
