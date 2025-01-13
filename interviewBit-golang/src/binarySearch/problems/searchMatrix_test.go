package bsearch

import (
	"fmt"
	"testing"
)

type searchMatrixTestCase struct {
	A        [][]int
	B        int
	Expected int
}

var searchMatrixTestCases = []searchMatrixTestCase{
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		3,
		1,
	},
	{
		[][]int{
			{5, 17, 100, 111},
			{119, 120, 127, 131},
		},
		3,
		0,
	},
}

func TestSearchMatrix(t *testing.T) {

	for idx, test := range searchMatrixTestCases {
		if output := SearchMatrix(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
