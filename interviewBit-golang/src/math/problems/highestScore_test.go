package mathProblems

import (
	"fmt"
	"testing"
)

type highestScoreTestCase struct {
	A        [][]string
	Expected int
}

var highestScoreTestCases = []highestScoreTestCase{
	{
		A:        [][]string{{"Bob", "80"}, {"Bob", "90"}, {"Alice", "90"}},
		Expected: 90,
	},
	{
		A:        [][]string{{"Bob", "80"}, {"Bob", "90"}, {"Alice", "90"}, {"Alice", "10"}},
		Expected: 85,
	},
}

func TestHighestScore(t *testing.T) {

	for idx, test := range highestScoreTestCases {
		if output := HighestScore(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
