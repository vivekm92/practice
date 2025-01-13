package stringProblems

import (
	"fmt"
	"testing"
)

type lengthOfLastWordTestCase struct {
	A        string
	Expected int
}

var lengthOfLastWordTestCases = []lengthOfLastWordTestCase{
	{"hello world", 5},
	{" hello world ", 5},
	{" hello world         ", 5},
	{"InterviewBit", 12},
}

func TestLengthOfLastWord(t *testing.T) {
	for idx, test := range lengthOfLastWordTestCases {
		if output := LengthOfLastWord(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range lengthOfLastWordTestCases {
		if output := LengthOfLastWord1(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
