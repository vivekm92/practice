package bsearch

import (
	"fmt"
	"testing"
)

type smallerOrEqualElementsTestCase struct {
	A        []int
	B        int
	Expected int
}

var smallerOrEqualElementsTestCases = []smallerOrEqualElementsTestCase{
	{
		[]int{1, 3, 4, 4, 6},
		4,
		4,
	},
	{
		[]int{1, 2, 5, 5},
		3,
		2,
	},
}

func TestSmallerOrEqualElements(t *testing.T) {

	for idx, test := range smallerOrEqualElementsTestCases {
		if output := SmallerOrEqualElements(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
