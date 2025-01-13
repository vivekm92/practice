package stringProblems

import (
	"fmt"
	"testing"
)

type minimumParanthesesTestCase struct {
	A        string
	Expected int
}

var minimumParanthesesTestCases = []minimumParanthesesTestCase{
	{")))()", 3},
	{")))(((", 6},
	{"(((((())))))", 0},
	{"())", 1},
	{"(((", 3},
}

func TestMinimumParantheses(t *testing.T) {
	for idx, test := range minimumParanthesesTestCases {
		if output := MinimumParantheses(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
