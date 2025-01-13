package mathProblems

import (
	"fmt"
	"testing"
)

type nextSimilarNumberTestCase struct {
	A        string
	Expected string
}

var nextSimilarNumberTestCases = []nextSimilarNumberTestCase{
	{"892795", "892957"},
	{"4321", "-1"},
	{"218765", "251678"},
}

func TestNextSimilarNumber(t *testing.T) {

	for idx, test := range nextSimilarNumberTestCases {
		if output := NextSimilarNumber(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
