package mathProblems

import (
	"fmt"
	"testing"
)

type addBitwiseTestCase struct {
	A        int
	B        int
	Expected int
}

var addBitwiseTestCases = []addBitwiseTestCase{
	{10, 12, 22},
	{-1, -4, -5},
	{0, 0, 0},
}

func TestAddBitwise(t *testing.T) {

	for idx, test := range addBitwiseTestCases {
		if output := AddBitwise(test.A, test.B); test.Expected != output {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
