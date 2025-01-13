package arrayProblems

import (
	"fmt"
	"testing"
)

type maximumGapTestCase struct {
	A        []int
	Expected int
}

var maximumGapTestCases = []maximumGapTestCase{
	{[]int{1, 10}, 9},
	{[]int{1, 1, 2}, 1},
	{[]int{5}, 0},
	{[]int{10, 1, 5}, 5},
	{[]int{4, 1, 2, 5, 3}, 1},
}

func TestMaximumGap(t *testing.T) {

	for idx, test := range maximumGapTestCases {
		if output := MaximumGap(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestMaximumGap1(t *testing.T) {
	for idx, test := range maximumGapTestCases {
		if output := MaximumGap1(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
