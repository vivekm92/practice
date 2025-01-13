package bitManipulation

import (
	"fmt"
	"testing"
)

type countTrailingZerosTestCase struct {
	A        int
	Expected int
}

var countTrailingZerosTestCases = []countTrailingZerosTestCase{
	{18, 1},
	{3, 0},
	{8, 3},
}

func TestCountTrailingZeros(t *testing.T) {
	for idx, test := range countTrailingZerosTestCases {
		if output := CountTrailingZeros(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
