package stringProblems

import (
	"fmt"
	"testing"
)

type serializeTestCase struct {
	A        []string
	Expected string
}

var serializeTestCases = []serializeTestCase{
	{[]string{"scaler", "academy"}, "scaler6~academy7~"},
	{[]string{"interviewbit"}, "interviewbit12~"},
}

func TestSerialize(t *testing.T) {
	for idx, test := range serializeTestCases {
		if output := Serialize(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
