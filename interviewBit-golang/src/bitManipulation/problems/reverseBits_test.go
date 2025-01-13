package bitManipulation

import (
	"fmt"
	"testing"
)

type reverseBitsTestCase struct {
	A        uint32
	Expected uint32
}

var reverseBitsTestCases = []reverseBitsTestCase{
	{0, 0},
	{3, 3221225472},
}

func TestReverseBits(t *testing.T) {
	for idx, test := range reverseBitsTestCases {
		if output := ReverseBits(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
