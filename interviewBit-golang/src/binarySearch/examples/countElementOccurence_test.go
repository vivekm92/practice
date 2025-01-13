package bsearch

import (
	"fmt"
	"testing"
)

type findElementCountTestCase struct {
	A        []int
	B        int
	Expected int
}

var occurenceOfEachNumnberTestCases = []findElementCountTestCase{
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		2,
	},
	{
		[]int{1},
		1,
		1,
	},
}

func TestFindElementCount(t *testing.T) {

	for idx, test := range occurenceOfEachNumnberTestCases {
		if output := FindElementCount(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
