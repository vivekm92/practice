package arrayProblems

import (
	"fmt"
	"testing"
)

type maxMinTestCase struct {
	A        []int
	Expected int
}

var maxMinTestCases = []maxMinTestCase{
	{
		[]int{-2, 1, -4, 5, 3},
		1,
	},
	{
		[]int{1, 3, 4, 1},
		5,
	},
}

func TestMaxMin(t *testing.T) {

	for idx, test := range maxMinTestCases {
		if output := MaxMin(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
