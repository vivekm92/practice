package bsearch

import (
	"fmt"
	"testing"
)

type paintTestCase struct {
	A        int
	B        int
	C        []int
	Expected int
}

var paintTestCases = []paintTestCase{
	{2, 5, []int{1, 10}, 50},
	{10, 1, []int{1, 8, 11, 3}, 11},
}

func TestPaint(t *testing.T) {
	for idx, test := range paintTestCases {
		if output := Paint(test.A, test.B, test.C); output != test.Expected {
			fmt.Println(test.A, test.B, test.C, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
