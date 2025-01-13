package bsearch

import (
	"fmt"
	"testing"
)

type findPeakTestCase struct {
	A        []int
	Expected int
}

var findPeakTestCases = []findPeakTestCase{
	{[]int{1, 2, 3, 4, 5}, 5},
	{[]int{5, 4, 3, 2, 1}, 5},
	{[]int{5, 71, 100, 11}, 100},
}

func TestFindPeak(t *testing.T) {

	for idx, test := range findPeakTestCases {
		if output := FindPeak(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestFindPeak1(t *testing.T) {

	for idx, test := range findPeakTestCases {
		if output := FindPeak1(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
