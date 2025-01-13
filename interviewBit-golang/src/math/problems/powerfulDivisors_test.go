package mathProblems

import (
	"reflect"
	"testing"
)

func TestPowerfulDivisors(t *testing.T) {
	tests := []struct {
		A        []int
		Expected []int
	}{
		{[]int{1, 4}, []int{1, 3}},
		{[]int{5, 10}, []int{4, 8}},
	}

	for _, test := range tests {
		result := PowerfulDivisors(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("PowerfulDivisors(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
