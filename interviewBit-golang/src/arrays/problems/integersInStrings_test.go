package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type integersInStringsTestCase struct {
	A        string
	Expected []int
}

var integersInStringsTestCases = []integersInStringsTestCase{
	{"1,2,3,4,5,6", []int{1, 2, 3, 4, 5, 6}},
}

func TestIntegersInStrings(t *testing.T) {
	for idx, test := range integersInStringsTestCases {
		if output := IntegersInStrings(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
