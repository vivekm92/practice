package arrayProblems

import (
	"fmt"
	"testing"
)

type climbingStairsTestCase struct {
	A        []int
	Expected int
}

var climbingStairsTestCases = []climbingStairsTestCase{
	{[]int{1, 2, 1, 3}, 5},
	{[]int{1, 2, 3, 4}, 7},
}

func TestClimbingStairs(t *testing.T) {
	for idx, test := range climbingStairsTestCases {
		if output := ClimbingStairs(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
