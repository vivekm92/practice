package bitManipulation

import (
	"fmt"
	"testing"
)

type countNumberOfSetBitsTestCase struct {
	A        int
	Expected int
}

var countNumberOfSetBitsTestCases = []countNumberOfSetBitsTestCase{
	{3, 2},
	{1024, 1},
	{0, 0},
	{255, 8}, // 2^7 + 2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 2^1 + 2^0
	{13, 3},
}

func TestCountNumberOfSetBits(t *testing.T) {
	for idx, test := range countNumberOfSetBitsTestCases {
		if output := CountNumberOfSetBits(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
