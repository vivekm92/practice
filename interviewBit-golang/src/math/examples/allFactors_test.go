package mathExamples

import (
	"fmt"
	"reflect"
	"testing"
)

type allFactorsTestCase struct {
	A        int
	Expected []int
}

var allFactorsTestCases = []allFactorsTestCase{
	{6, []int{1, 2, 3, 6}},
	{11, []int{1, 11}},
}

func TestAllFactors(t *testing.T) {
	for idx, test := range allFactorsTestCases {
		if output := AllFactors(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
