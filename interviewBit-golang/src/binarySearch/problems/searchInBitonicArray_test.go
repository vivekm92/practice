package bsearch

import (
	"fmt"
	"testing"
)

type searchInBitnoicArrayTestCase struct {
	A        []int
	B        int
	Expected int
}

var searchInBitnoicArrayTestCases = []searchInBitnoicArrayTestCase{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
		12,
		18,
	},
	{
		[]int{3, 9, 10, 20, 17, 5, 1},
		20,
		3,
	},
	{
		[]int{5, 6, 7, 8, 9, 10, 3, 2, 1},
		30,
		-1,
	},
}

func TestSearchInBitonicArray(t *testing.T) {

	for idx, test := range searchInBitnoicArrayTestCases {
		if output := SearchInBitnoicArray(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
