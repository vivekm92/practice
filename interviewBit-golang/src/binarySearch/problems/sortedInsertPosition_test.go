package bsearch

import (
	"fmt"
	"testing"
)

type searchInsertTestCase struct {
	A        []int
	B        int
	Expected int
}

var searchInsertTestCases = []searchInsertTestCase{
	{
		[]int{1, 3, 5, 6},
		5,
		2,
	},
	{
		[]int{1, 3, 5, 6},
		2,
		1,
	},
}

func TestSearchInsert(t *testing.T) {

	for idx, test := range searchInsertTestCases {
		if output := SearchInsert(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
