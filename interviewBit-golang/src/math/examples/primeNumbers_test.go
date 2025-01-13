package mathExamples

import (
	"fmt"
	"reflect"
	"testing"
)

type primeNumbersTestCase struct {
	A        int
	Expected []int
}

var primeNumbersTestCases = []primeNumbersTestCase{
	{11, []int{2, 3, 5, 7, 11}},
}

func TestPrimeNumbers(t *testing.T) {
	for idx, test := range primeNumbersTestCases {
		if output := PrimeNumbers(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestSieve(t *testing.T) {
	for idx, test := range primeNumbersTestCases {
		if output := Sieve(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
