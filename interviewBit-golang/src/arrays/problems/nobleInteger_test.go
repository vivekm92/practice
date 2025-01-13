package arrayProblems

import (
	"fmt"
	"testing"
)

type nobleIntegerTestCase struct {
	A        []int
	Expected int
}

var nobleIntegerTestCases = []nobleIntegerTestCase{
	{
		[]int{3, 2, 1, 3},
		1,
	},
	{
		[]int{1, 1, 3, 3},
		-1,
	},
}

func TestNobleIntegers(t *testing.T) {

	for idx, test := range nobleIntegerTestCases {
		if output := NobleInteger(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range nobleIntegerTestCases {
		if output := NobleIntegers(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
