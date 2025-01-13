package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type findPermutationTestCase struct {
	A        string
	B        int
	Expected []int
}

var findPermutationTestCases = []findPermutationTestCase{
	{
		"IIIII",
		6,
		[]int{1, 2, 3, 4, 5, 6},
	},
	{
		"DDD",
		4,
		[]int{4, 3, 2, 1},
	},
	{
		"DIDD",
		5,
		[]int{5, 1, 4, 3, 2},
	},
}

func TestFindPermutation(t *testing.T) {

	for idx, test := range findPermutationTestCases {
		if output := FindPermutation(test.A, test.B); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
