package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type occurenceOfEachNumnberTestCase struct {
	A        []int
	Expected []int
}

var occurenceOfEachNumnberTestCases = []occurenceOfEachNumnberTestCase{
	{
		[]int{1, 2, 3},
		[]int{1, 1, 1},
	},
	{
		[]int{4, 3, 3},
		[]int{2, 1},
	},
}

func TestOccurenceOfEachNumber(t *testing.T) {

	for idx, test := range occurenceOfEachNumnberTestCases {
		if output := OccurenceOfEachNumnber(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)

		}
	}
}
