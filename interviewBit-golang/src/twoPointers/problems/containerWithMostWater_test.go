package twoPointers

import (
	"fmt"
	"testing"
)

type maxAreaTestCase struct {
	A        []int
	Expected int
}

var maxAreaTestCases = []maxAreaTestCase{
	{[]int{1, 5, 4, 3}, 6},
}

func TestMaxArea(t *testing.T) {
	for idx, test := range maxAreaTestCases {
		if output := MaxArea(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
