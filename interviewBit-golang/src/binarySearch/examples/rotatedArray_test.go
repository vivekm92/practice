package bsearch

import (
	"fmt"
	"testing"
)

type findMinTestCase struct {
	A        []int
	Expected int
}

var findMinTestCases = []findMinTestCase{
	{[]int{4, 5, 6, 0, 1, 2, 3}, 0},
	{[]int{5, 6, 0, 1, 2, 3, 4}, 0},
	{[]int{6, 0, 1, 2, 3, 4, 5}, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6}, 0},
}

func TestFindMin(t *testing.T) {
	for idx, test := range findMinTestCases {
		if output := FindMin(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range findMinTestCases {
		if output := FindMin2(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
