package twoPointers

import (
	"fmt"
	"reflect"
	"testing"
)

type mergeTestCase struct {
	A        []int
	B        []int
	Expected []int
}

var mergeTestCases = []mergeTestCase{
	{[]int{1, 5, 8}, []int{6, 9}, []int{1, 5, 6, 8, 9}},
}

func TestMerge(t *testing.T) {
	for idx, test := range mergeTestCases {
		if output := Merge(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
