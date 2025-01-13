package bsearch

import (
	"fmt"
	"testing"
)

type allocateBooksTestCase struct {
	A        []int
	B        int
	Expected int
}

var allocateBooksTestCases = []allocateBooksTestCase{
	{[]int{31, 14, 19, 75}, 12, -1},
	{[]int{12, 34, 67, 90}, 2, 113},
	{[]int{5, 17, 100, 11}, 4, 100},
}

func TestAllocateBooks(t *testing.T) {
	for idx, test := range allocateBooksTestCases {
		if output := AllocateBooks(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
