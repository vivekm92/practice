package stacksAndQueues

import (
	"fmt"
	"testing"
)

type balancedParanthesesTestCase struct {
	A        string
	Expected int
}

var balancedParanthesesTestCases = []balancedParanthesesTestCase{
	{"(()())", 1},
	{"(()", 0},
}

func TestBalancedParantheses(t *testing.T) {
	for idx, test := range balancedParanthesesTestCases {
		if output := BalancedParantheses(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range balancedParanthesesTestCases {
		if output := BalancedParantheses1(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
