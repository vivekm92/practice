package arrayProblems

import (
	"fmt"
	"testing"
)

type minimumLightsToActivateTestCase struct {
	A        []int
	B        int
	Expected int
}

var minimumLightsToActivateTestCases = []minimumLightsToActivateTestCase{
	{
		[]int{0, 0, 1, 1, 1, 0, 0, 1},
		3,
		2,
	},
	{
		[]int{0, 0, 0, 1, 0},
		3,
		-1,
	},
}

func TestMinimumLightsToActivate(t *testing.T) {

	for idx, test := range minimumLightsToActivateTestCases {
		if output := MinimumLightsToActivate(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
