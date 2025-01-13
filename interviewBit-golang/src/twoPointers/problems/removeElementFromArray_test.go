package twoPointers

import (
	"fmt"
	"testing"
)

type removeElementsTestCase struct {
	A        []int
	B        int
	Expected int
}

var removeElementsTestCases = []removeElementsTestCase{
	{[]int{1, 2, 3, 4, 5, 7, 7, 7, 2}, 2, 7},
}

func TestRemoveElememnts(t *testing.T) {

	for idx, test := range removeElementsTestCases {
		if output := RemoveElements(&test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		} else {
			// check if test.B is present in [0, output - 1]
			for i := 0; i < output; i++ {
				if test.A[i] == test.B {
					t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
				}
			}
		}
	}
}
