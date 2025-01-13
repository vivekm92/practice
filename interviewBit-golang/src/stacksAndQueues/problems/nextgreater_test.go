package stacksAndQueues

import (
	"reflect"
	"testing"
)

func TestNextGreater(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, -1}},
	}

	for _, test := range tests {
		result := NextGreater(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("NextGreater(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
