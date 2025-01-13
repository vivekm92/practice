package twoPointers

import (
	"fmt"
	"testing"
)

type threeSumClosestTestCase struct {
	A        []int
	B        int
	Expected int
}

var threeSumClosestTestCases = []threeSumClosestTestCase{
	{[]int{-1, 2, 1, -4}, 1, 2},
	{[]int{1, 2, 3}, 6, 6},
}

func TestThreeSumClosest(t *testing.T) {
	for idx, test := range threeSumClosestTestCases {
		if output := ThreeSumClosest(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
