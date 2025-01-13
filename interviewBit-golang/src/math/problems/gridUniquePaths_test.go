package mathProblems

import (
	"fmt"
	"testing"
)

type uniquePathsTestCase struct {
	A        int
	B        int
	Expected int
}

var uniquePathsTestCases = []uniquePathsTestCase{
	{100, 1, 1},
	{2, 2, 2},
	{3, 6, 21},
}

func TestUniquePaths(t *testing.T) {
	for idx, test := range uniquePathsTestCases {
		if output := UniquePaths(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
