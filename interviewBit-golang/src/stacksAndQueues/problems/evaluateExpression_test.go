package stacksAndQueues

import (
	"fmt"
	"testing"
)

type evalRPNTestCase struct {
	A        []string
	Expected int
}

var evalRPNTestCases = []evalRPNTestCase{
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
}

func TestEvalRPN(t *testing.T) {
	for idx, test := range evalRPNTestCases {
		if output := EvalRPN(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
