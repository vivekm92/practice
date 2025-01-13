package bsearch

import (
	"fmt"
	"testing"
)

type sqrtTestCase struct {
	A        int
	Expected int
}

var sqrtTestCases = []sqrtTestCase{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 2},
	{9, 3},
	{10, 3},
	{24, 4},
}

func TestSqrt(t *testing.T) {

	for idx, test := range sqrtTestCases {
		if output := Sqrt(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
