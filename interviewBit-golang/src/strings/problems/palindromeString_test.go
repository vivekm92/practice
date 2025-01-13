package stringProblems

import (
	"fmt"
	"testing"
)

type isValidCharTestCase struct {
	A        byte
	Expected bool
}

var isValidCharTestCases = []isValidCharTestCase{
	{'a', true},
	{'$', false},
	{'A', true},
	{'8', true},
}

func TestIsValidChar(t *testing.T) {
	for idx, test := range isValidCharTestCases {
		if output := IsValidChar(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

type isEqualCharTestCase struct {
	A        byte
	B        byte
	Expected bool
}

var isEqualCharTestCases = []isEqualCharTestCase{
	{'1', '8', false},
	{'a', 'A', true},
	{'A', ' ', false},
	{'a', 'a', true},
	{'A', 'A', true},
}

func TestIsEqualChart(t *testing.T) {
	for idx, test := range isEqualCharTestCases {
		if output := IsEqualChar(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

type isPalindromeTestCase struct {
	A        string
	Expected int
}

var isPalindromeTestCases = []isPalindromeTestCase{
	{"A man, a plan, a canal: Panama", 1},
}

func TestIsPalindrome(t *testing.T) {
	for idx, test := range isPalindromeTestCases {
		if output := IsPalindrome(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
