package twoPointers

import (
	"fmt"
	"reflect"
	"testing"
)

type intersectTestCase struct {
	A        []int
	B        []int
	Expected []int
}

var intersectTestCases = []intersectTestCase{
	{[]int{1, 2, 3, 3, 4, 5, 6}, []int{3, 3, 5}, []int{3, 3, 5}},
	{[]int{1, 2, 3, 3, 4, 5, 6}, []int{3, 5}, []int{3, 5}},
	{[]int{1, 2, 3, 4}, []int{7, 8, 9}, []int{}},
}

func TestIntersect(t *testing.T) {
	for idx, test := range intersectTestCases {
		if output := Intersect(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
