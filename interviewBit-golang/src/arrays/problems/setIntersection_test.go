package arrayProblems

import (
	"fmt"
	"testing"
)

type setIntersectionTestCase struct {
	A        [][]int
	Expected int
}

var setIntersectionTestCases = []setIntersectionTestCase{
	{
		[][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
		3,
	},
	{
		[][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
		5,
	},
}

func TestSetIntersection(t *testing.T) {

	for idx, test := range setIntersectionTestCases {
		if output := SetIntersection(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
