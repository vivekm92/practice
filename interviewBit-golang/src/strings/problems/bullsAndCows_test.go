package stringProblems

import (
	"fmt"
	"testing"
)

type bullsAndCowsTestCase struct {
	A        string
	B        string
	Expected string
}

var bullsAndCowsTestCases = []bullsAndCowsTestCase{
	{"1807", "7810", "1A3B"},
	{"1123", "0111", "1A1B"},
	{"6020943525972", "7157627311068", "0A6B"},
}

func TestBullsAndCows(t *testing.T) {
	for idx, test := range bullsAndCowsTestCases {
		if output := BullsAndCows(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
