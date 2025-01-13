package arrayProblems

import (
	"fmt"
	"testing"
)

type maximumSumSquareSubmatrixTestCase struct {
	A        [][]int
	B        int
	Expected int
}

var maximumSumSquareSubmatrixTestCases = []maximumSumSquareSubmatrixTestCase{
	{
		[][]int{
			{1, 1, 1, 1, 1},
			{2, 2, 2, 2, 2},
			{3, 8, 6, 7, 3},
			{4, 4, 4, 4, 4},
			{5, 5, 5, 5, 5},
		}, 3, 48,
	},
	{
		[][]int{
			{2, 2},
			{2, 2},
		},
		2,
		8,
	},
}

func TestMaximumSum(t *testing.T) {

	for idx, test := range maximumSumSquareSubmatrixTestCases {
		if output := MaximumSum(test.A, test.B); test.Expected != output {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
