package bsearch

import (
	"fmt"
	"testing"
)

type woodCuttingMadeEasyTestCase struct {
	A        []int
	B        int
	Expected int
}

var woodCuttingMadeEasyTestCases = []woodCuttingMadeEasyTestCase{
	{
		[]int{20, 15, 10, 17},
		7,
		15,
	},
	{
		[]int{4, 42, 40, 26, 46},
		20,
		36,
	},
}

func TestMaximumSawBladeHeight(t *testing.T) {

	for idx, test := range woodCuttingMadeEasyTestCases {
		if output := MaximumSawBladeHeight(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
