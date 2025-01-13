package stacksAndQueues

import (
	"fmt"
	"testing"
)

type nonRepeatingCharacters1TestCase struct {
	A        string
	Expected string
}

var nonRepeatingCharacters1TestCases = []nonRepeatingCharacters1TestCase{
	{"abadbc", "aabbdd"},
	{"abcabc", "aaabc#"},
}

func TestNonRepeatingCharacters1(t *testing.T) {
	for idx, test := range nonRepeatingCharacters1TestCases {
		if output := NonRepeatingCharacters(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range nonRepeatingCharacters1TestCases {
		if output := NonRepeatingCharacters1(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
