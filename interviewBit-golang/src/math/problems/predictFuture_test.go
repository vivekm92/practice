package mathProblems

import (
	"reflect"
	"testing"
)

func TestPredictFuture(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		C        int
		Expected []int
	}{
		{1, 3, 1, []int{1, 7}},
		{1, 3, 2, []int{5, 15}},
		{1, 3, 3, []int{5, 35}},
	}

	for _, test := range tests {
		result := PredictFuture(test.A, test.B, test.C)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("PredictFuture(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
