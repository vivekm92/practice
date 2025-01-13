package twoPointers

import (
	"fmt"
	"testing"
)

type diffPossibleTestCase struct {
	A        []int
	B        int
	Expected int
}

var diffPossibleTestCases = []diffPossibleTestCase{
	{[]int{1, 3, 5}, 4, 1},
	{[]int{1, 2, 2, 4}, 0, 1},
	{[]int{0}, 0, 0},
}

func TestDiffPossible(t *testing.T) {
	for idx, test := range diffPossibleTestCases {
		if output := DiffPossible(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
