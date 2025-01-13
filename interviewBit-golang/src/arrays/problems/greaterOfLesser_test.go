package arrayProblems

import (
	"fmt"
	"testing"
)

type greaterOrLesserTestCase struct {
	A        []int
	B        []int
	C        int
	Expected int
}

var greaterOrLesserTestCases = []greaterOrLesserTestCase{
	{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		4,
		0,
	},
	{
		[]int{1, 10, 100},
		[]int{9, 9, 9},
		50,
		3,
	},
}

func TestCountGreaterOrLesser(t *testing.T) {

	for idx, test := range greaterOrLesserTestCases {
		if output := CountGreaterOrLesser(test.A, test.B, test.C); output != test.Expected {
			fmt.Println(test.A, test.B, test.C, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
