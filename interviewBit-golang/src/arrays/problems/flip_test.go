package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type flipTestCase struct {
	A        string
	Expected []int
}

var flipTestCases = []flipTestCase{
	{"111", []int{}},
	{"010", []int{1, 1}},
	{"10100010111", []int{2, 6}},
}

func TestFlip(t *testing.T) {

	for idx, test := range flipTestCases {
		if output := Flip(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
