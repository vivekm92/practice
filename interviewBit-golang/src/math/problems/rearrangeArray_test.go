package mathProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type arrangeTestCase struct {
	A        []int
	Expected []int
}

var arrangeTestCases = []arrangeTestCase{
	{[]int{3, 2, 1, 4, 0, 5}, []int{4, 1, 2, 0, 3, 5}},
}

func TestArrange(t *testing.T) {

	for idx, test := range arrangeTestCases {
		if output := Arrange(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
