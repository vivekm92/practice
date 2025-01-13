package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type antiDiagonalTestCase struct {
	A        [][]int
	Expected [][]int
}

var antiDiagonalTestCases = []antiDiagonalTestCase{
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1}, {2, 4}, {3, 5, 7}, {6, 8}, {9}}},
	{[][]int{{1, 2}, {3, 4}}, [][]int{{1}, {2, 3}, {4}}},
}

func TestAntiDiagonals(t *testing.T) {

	for idx, test := range antiDiagonalTestCases {
		if output := AntiDiagonal(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
