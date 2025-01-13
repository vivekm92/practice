package bsearch

import (
	"fmt"
	"testing"
)

type powTestCase struct {
	A        int
	B        int
	C        int
	Expected int
}

var powTestCases = []powTestCase{
	{0, 0, 1, 0}, {-1, 1, 20, 19}, {71045970, 41535484, 64735492, 20805472},
}

func TestPow(t *testing.T) {

	for idx, test := range powTestCases {
		if output := Pow(test.A, test.B, test.C); output != test.Expected {
			fmt.Println(test.A, test.B, test.C, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
