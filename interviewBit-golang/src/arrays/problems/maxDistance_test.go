package arrayProblems

import (
	"fmt"
	"testing"
)

type maxDistanceTestCase struct {
	A        []int
	Expected int
}

var maxDistanceTestCases = []maxDistanceTestCase{
	{
		[]int{3, 5, 4, 2},
		2,
	},
}

func TestMaxDistance(t *testing.T) {

	for idx, test := range maxDistanceTestCases {
		if output := MaxDistance(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
