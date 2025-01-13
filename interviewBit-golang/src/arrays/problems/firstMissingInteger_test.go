package arrayProblems

import (
	"fmt"
	"testing"
)

type firstMissingPositiveTestCase struct {
	A        []int
	Expected int
}

var firstMissingPositiveTestCases = []firstMissingPositiveTestCase{
	{[]int{1, 2, 0}, 3},
	{[]int{3, 4, -1, 1}, 2},
	{[]int{-8, -7, -6}, 1},
	{[]int{-5}, 1},
}

func TestFirstMissingPositive(t *testing.T) {

	for idx, test := range firstMissingPositiveTestCases {
		if output := FirstMissingPositive(test.A); test.Expected != output {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
