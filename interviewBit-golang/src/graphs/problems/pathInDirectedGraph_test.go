package graphs

import (
	"fmt"
	"testing"
)

type pathInDirectedGraphTestCase struct {
	A        int
	B        [][]int
	Expected int
}

var pathInDirectedGraphTestCases = []pathInDirectedGraphTestCase{
	{5, [][]int{{1, 2}, {4, 1}, {2, 4}, {3, 4}, {5, 2}, {1, 3}}, 0},
	{5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 1},
}

func TestPathInDirectedGraph(t *testing.T) {
	for idx, test := range pathInDirectedGraphTestCases {
		if output := PathInDirectedGraph(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestPathInDirectedGraph1(t *testing.T) {
	for idx, test := range pathInDirectedGraphTestCases {
		if output := PathInDirectedGraph1(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
