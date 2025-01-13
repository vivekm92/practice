package twoPointers

import (
	"fmt"
	"testing"
)

type isPalindromeTestCase struct {
	A        int
	Expected bool
}

var isPalindromeTestCases = []isPalindromeTestCase{
	{101, true},
	{55, true},
	{1001, true},
	{27, false},
	{1, true},
}

func TestIsPalindromeNumber(t *testing.T) {
	for idx, test := range isPalindromeTestCases {
		if output := IsPalindrome(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

type countPalindromeNumbersTestCase struct {
	A        int
	B        int
	C        int
	Expected int
}

var countPalindromeNumbersTestCases = []countPalindromeNumbersTestCase{
	{88, 104, 3, 2},
	{88, 104, 10, 2},
	{1, 10, 10, 9},
}

func TestCountPalindromeNumbers(t *testing.T) {
	for idx, test := range countPalindromeNumbersTestCases {
		if output := CountPalindromeNumbers(test.A, test.B, test.C); test.Expected != output {
			fmt.Println(test.A, test.B, test.C, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
