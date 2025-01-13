package mathExamples

import (
	"fmt"
	"testing"
)

type findDigitsInBinaryTestCase struct {
	A        int
	Expected string
}

var findDigitsInBinaryTestCases = []findDigitsInBinaryTestCase{
	{0, "0"},
	{1, "1"},
	{2, "10"},
	{5, "101"},
	{13, "1101"},
}

func TestFindDigitsInBinary(t *testing.T) {
	for idx, test := range findDigitsInBinaryTestCases {
		if output := FindDigitsInBinary(test.A); test.Expected != output {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
