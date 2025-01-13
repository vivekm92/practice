package stringProblems

import (
	"fmt"
	"testing"
)

type zigZagStringTestCase struct {
	A        string
	B        int
	Expected string
}

var zigZagStringTestCases = []zigZagStringTestCase{
	{"ABCD", 2, "ACBD"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"ABCDEF", 6, "ABCDEF"},
	{"ABCDEF", 1, "ABCDEF"},
	{"ABCDEFGHIJKLMNOPQRS", 6, "AKBJLCIMSDHNREGOQFP"},
}

func TestZigZagString(t *testing.T) {
	for idx, test := range zigZagStringTestCases {
		if output := ZigZagString(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestZigZagString1(t *testing.T) {
	for idx, test := range zigZagStringTestCases {
		if output := ZigZagString1(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
