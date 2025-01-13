package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type rotateArrayTestCase struct {
	A        []int
	B        int
	Expected []int
}

var rotateArrayTestCases = []rotateArrayTestCase{
	{[]int{1, 2, 3, 4, 5, 6}, 1, []int{2, 3, 4, 5, 6, 1}},
}

func TestRotateArray(t *testing.T) {
	for idx, test := range rotateArrayTestCases {
		if output := RotateArray(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
