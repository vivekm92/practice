package twoPointers

import (
	"fmt"
	"testing"
)

type pairsWithGivenDifferenceTestCase struct {
	A        []int
	B        int
	Expected int
}

var pairsWithGivenDifferenceTestCases = []pairsWithGivenDifferenceTestCase{
	{[]int{5, 10, 3, 2, 50, 80}, 78, 1},
	{[]int{-10, 20}, 30, 1},
	{[]int{2, 31, 5, 3, 4}, 0, 0},
	{[]int{2, 31, 5, 2, 4}, 0, 1},
}

func TestPairsWithGivenDifference(t *testing.T) {
	for idx, test := range pairsWithGivenDifferenceTestCases {
		if output := PairsWithGivenDifference(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
