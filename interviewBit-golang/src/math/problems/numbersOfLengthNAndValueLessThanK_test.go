package mathProblems

import (
	"fmt"
	"testing"
)

type numbersOfLengthNAndValueLessThanKTestCase struct {
	A        []int
	B        int
	C        int
	Expected int
}

var numbersOfLengthNAndValueLessThanKTestCases = []numbersOfLengthNAndValueLessThanKTestCase{
	{[]int{0, 1, 5}, 1, 2, 2},
	{[]int{0, 1, 2, 5}, 2, 21, 5},
}

func TestNumbersOfLengthNAndValueLessThanK(t *testing.T) {

	for idx, test := range numbersOfLengthNAndValueLessThanKTestCases {
		if output := NumbersOfLengthNAndValueLessThanK(test.A, test.B, test.C); output != test.Expected {
			fmt.Println(test.A, test.B, test.C, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
