package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type subUnsortTestCase struct {
	A        []int
	Expected []int
}

var subUnsortTestCases = []subUnsortTestCase{
	{
		[]int{4, 15, 4, 4, 15, 18, 20},
		[]int{1, 3},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{-1},
	},
	{
		[]int{1, 3, 2, 4, 5},
		[]int{1, 2},
	},
	{
		[]int{1, 2, 10, 12, 15, 6, 7, 8, 9, 20},
		[]int{2, 8},
	},
}

func TestSubUnsort(t *testing.T) {

	for idx, test := range subUnsortTestCases {
		if output := SubUnsort(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
