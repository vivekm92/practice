package twoPointers

import (
	"fmt"
	"reflect"
	"testing"
)

type sortColorTestCase struct {
	A        []int
	Expected []int
}

var sortColorTestCases = []sortColorTestCase{
	{[]int{0, 1, 2, 0, 1, 2}, []int{0, 0, 1, 1, 2, 2}},
	{[]int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
	{[]int{0, 2, 0, 2}, []int{0, 0, 2, 2}},
	{[]int{0, 1, 0, 1, 0, 1, 2, 0, 1, 2, 1, 2, 0, 1, 1, 2, 1}, []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2}},
	{[]int{0}, []int{0}},
}

func TestSortColor(t *testing.T) {

	for idx, test := range sortColorTestCases {
		if output := SortColor(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range sortColorTestCases {
		if output := SortColor2(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range sortColorTestCases {
		if output := SortColor3(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
