package stacksAndQueues

import (
	"fmt"
	"reflect"
	"testing"
)

type prevSmallerTestCase struct {
	A        []int
	Expected []int
}

var prevSmallerTestCases = []prevSmallerTestCase{
	{[]int{39, 27, 11, 4, 24, 32, 32, 1}, []int{-1, -1, -1, -1, 4, 24, 24, -1}},
	{[]int{4, 5, 2, 10, 8}, []int{-1, 4, -1, 2, 2}},
	{[]int{3, 2, 1}, []int{-1, -1, -1}},
}

func TestPrevSmaller(t *testing.T) {
	for idx, test := range prevSmallerTestCases {
		if output := PrevSmaller(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
