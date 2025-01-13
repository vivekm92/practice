package bsearch

import (
	"fmt"
	"testing"
)

type findLeastCapacityTestCase struct {
	A        []int
	B        int
	Expected int
}

var findLeastCapacityTestCases = []findLeastCapacityTestCase{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
	{[]int{3, 2, 2, 4, 1, 4}, 3, 6},
}

func TestFindLeastCapacity(t *testing.T) {

	for idx, test := range findLeastCapacityTestCases {
		if output := FindLeastCapacity(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
