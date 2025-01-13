package mathExamples

import (
	"fmt"
	"testing"
)

type sampleQuestionTestCase struct {
	A        int
	B        int
	Expected int
}

var sampleQuestionTestCases = []sampleQuestionTestCase{
	{2, 3, 5},
	{0, 10000000, 0},
}

func TestSampleQuestion(t *testing.T) {
	for idx, test := range sampleQuestionTestCases {
		if output := SampleQuestion(test.A, test.B); test.Expected != output {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
