package mathExamples

import (
	"fmt"
	"testing"
)

type isPrimeTestCase struct {
	A        int
	Expected int
}

var isPrimeTestCases = []isPrimeTestCase{
	{1, 0},
	{2, 1},
	{49, 0},
}

func TestIsPrime(t *testing.T) {
	for idx, test := range isPrimeTestCases {
		if output := IsPrime(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
