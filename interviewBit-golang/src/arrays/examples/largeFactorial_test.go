package arrayExamples

import "testing"

type largeFactorialTestCase struct {
	A        int
	Expected string
}

var largeFactorialTestCases = []largeFactorialTestCase{
	{5, "120"},
	{100, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
}

func TestLargeFactorial(t *testing.T) {
	for idx, test := range largeFactorialTestCases {
		if output := LargeFactorial(test.A); output != test.Expected {
			t.Errorf(" failed : %v, output : %v", idx, output)
		}
	}
}
