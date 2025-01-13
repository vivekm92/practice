package stringProblems

import (
	"fmt"
	"testing"
)

type salutesTestCase struct {
	A        string
	Expected int64
}

var salutesTestCases = []salutesTestCase{
	{">>><<<", 9},
	{"<>", 0},
	{"><", 1},
}

func TestCountSalutes(t *testing.T) {
	for idx, test := range salutesTestCases {
		if output := CountSalutes(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
