package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type nextPermutationTestCase struct {
	A        []int
	Expected []int
}

var nextPermutationTestCases = []nextPermutationTestCase{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	},
	{
		[]int{3, 2, 1},
		[]int{1, 2, 3},
	},
	{
		[]int{1, 1, 5},
		[]int{1, 5, 1},
	},
	{
		[]int{20, 50, 113},
		[]int{20, 113, 50},
	},
}

func TestNextPermutation(t *testing.T) {

	for idx, test := range nextPermutationTestCases {
		if output := NextPermutation(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
