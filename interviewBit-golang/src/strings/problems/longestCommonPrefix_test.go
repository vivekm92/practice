package stringProblems

import (
	"fmt"
	"testing"
)

type longestCommonPrefixTestCase struct {
	A        []string
	Expected string
}

var longestCommonPrefixTestCases = []longestCommonPrefixTestCase{
	{[]string{"abcdefgh", "aefghijk", "abcefgh"}, "a"},
	{[]string{"abab", "ab", "abcd"}, "ab"},
	{[]string{"abc", "def"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for idx, test := range longestCommonPrefixTestCases {
		if output := LongestCommonPrefix(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
