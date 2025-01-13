package bitManipulation

import (
	"fmt"
	"testing"
)

type singleNumberTestCase struct {
	A        []int
	Expected int
}

var singleNumberTestCases = []singleNumberTestCase{
	{[]int{2, 3, 4, 1, 2, 3, 4}, 1},
	{[]int{1, 2, 3, 4, 1, 2, 3}, 4},
}

func TestSingleNumber(t *testing.T) {
	for idx, test := range singleNumberTestCases {
		if output := SingleNumber(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
