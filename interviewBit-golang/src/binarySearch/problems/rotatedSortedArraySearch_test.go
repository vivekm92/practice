package bsearch

import (
	"fmt"
	"testing"
)

type rotatedSortedArraySearch struct {
	A        []int
	B        int
	Expected int
}

var rotatedSortedArraySearchTestCases = []rotatedSortedArraySearch{
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 4, 0},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 9, -1},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 3, 7},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 5, 1},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 2, 6},
	{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 1, 5},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 0, 0},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 7, 7},
	{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 2, 2},
	{[]int{6, 0, 1, 2, 3, 4, 5}, 6, 0},
}

func TestRotatedSearch(t *testing.T) {
	for idx, test := range rotatedSortedArraySearchTestCases {
		if output := RotatedSearch(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
