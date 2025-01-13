package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type prettyPrintTestCase struct {
	A        int
	Expected [][]int
}

var prettyprintTestCases = []prettyPrintTestCase{
	{3, [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}},
	{4, [][]int{{4, 4, 4, 4, 4, 4, 4}, {4, 3, 3, 3, 3, 3, 4}, {4, 3, 2, 2, 2, 3, 4}, {4, 3, 2, 1, 2, 3, 4}, {4, 3, 2, 2, 2, 3, 4}, {4, 3, 3, 3, 3, 3, 4}, {4, 4, 4, 4, 4, 4, 4}}},
	{1, [][]int{{1}}},
	{2, [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
}

func TestPrettyPrint(t *testing.T) {

	for idx, test := range prettyprintTestCases {
		if output := PrettyPrint(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
