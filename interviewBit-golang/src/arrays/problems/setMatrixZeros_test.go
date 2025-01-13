package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type setZerosTestCase struct {
	A        [][]int
	Expected [][]int
}

var setZerosTestCases = []setZerosTestCase{
	{
		[][]int{
			{1, 0, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		[][]int{
			{0, 0, 0},
			{1, 0, 1},
			{1, 0, 1},
		},
	},
	{
		[][]int{
			{1, 0, 1},
			{1, 1, 1},
			{1, 0, 1},
		},
		[][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
	},
}

func TestSetZeros(t *testing.T) {
	for idx, test := range setZerosTestCases {
		if output := SetZeros(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
