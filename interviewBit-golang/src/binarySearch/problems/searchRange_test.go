package bsearch

import (
	"fmt"
	"reflect"
	"testing"
)

type searchRangeTestCase struct {
	A        []int
	B        int
	Expected []int
}

var searchRangeTestCases = []searchRangeTestCase{
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		[]int{5, 17, 100, 111},
		3,
		[]int{-1, -1},
	},
}

func TestSearchRange(t *testing.T) {

	for idx, test := range searchRangeTestCases {
		if output := SearchRange(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
