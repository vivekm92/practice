package bitManipulation

import (
	"fmt"
	"testing"
)

type minXorValueTestCase struct {
	A        []int
	Expected int
}

var minXorValueTestCases = []minXorValueTestCase{
	{[]int{0, 2, 5, 7}, 2},
	{[]int{0, 4, 7, 9}, 3},
}

func TestFindMinXor(t *testing.T) {
	for idx, test := range minXorValueTestCases {
		if output := FindMinXor(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
