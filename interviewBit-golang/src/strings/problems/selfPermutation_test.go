package stringProblems

import (
	"fmt"
	"testing"
)

type permuteStringsTestCase struct {
	A        string
	B        string
	Expected int
}

var permuteStringsTestCases = []permuteStringsTestCase{
	{"scaler", "relasc", 1},
	{"scaler", "interviewvit", 0},
	{"scaler", "srelasc", 0},
}

func TestPermuteStrings(t *testing.T) {
	for idx, test := range permuteStringsTestCases {
		if output := PermuteStrings(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
