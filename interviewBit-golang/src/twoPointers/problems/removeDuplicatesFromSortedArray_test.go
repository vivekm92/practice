package twoPointers

import (
	"fmt"
	"reflect"
	"testing"
)

type removeDuplicatesTestCase struct {
	A        []int
	Expected []int
}

var removeDuplicatesTestCases = []removeDuplicatesTestCase{
	{[]int{1, 1, 2}, []int{1, 2}},
	{[]int{1, 2, 2, 3, 3}, []int{1, 2, 3}},
}

func TestRemoveDuplicates(t *testing.T) {

	for idx, test := range removeDuplicatesTestCases {
		if output := RemoveDuplicates(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
