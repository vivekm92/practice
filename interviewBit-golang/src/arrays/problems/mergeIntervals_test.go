package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type mergeIntervalTestCase struct {
	A        []Interval
	B        Interval
	Expected []Interval
}

var mergeIntervalTestCases = []mergeIntervalTestCase{
	{
		[]Interval{{1, 3}, {5, 6}},
		Interval{8, 10},
		[]Interval{{1, 3}, {5, 6}, {8, 10}},
	},
	{
		[]Interval{{3, 4}, {5, 6}},
		Interval{1, 2},
		[]Interval{{1, 2}, {3, 4}, {5, 6}},
	},
	{
		[]Interval{{1, 3}, {5, 6}},
		Interval{2, 4},
		[]Interval{{1, 4}, {5, 6}},
	},
	{
		[]Interval{{1, 3}, {5, 6}},
		Interval{2, 5},
		[]Interval{{1, 6}},
	},
}

func TestMergeIntervals(t *testing.T) {

	for idx, test := range mergeIntervalTestCases {
		if output := MergeInterval(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
