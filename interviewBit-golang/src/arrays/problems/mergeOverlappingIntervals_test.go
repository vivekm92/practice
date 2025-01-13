package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type mergeOverlappingIntervalsTestCase struct {
	A        []Interval
	Expected []Interval
}

var mergeOverlappingIntervalsTestCases = []mergeOverlappingIntervalsTestCase{
	{
		[]Interval{},
		[]Interval{},
	},
	{
		[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		[]Interval{{1, 6}, {8, 10}, {15, 18}},
	},
}

func TestMergeOverlappingIntervals(t *testing.T) {

	for idx, test := range mergeOverlappingIntervalsTestCases {
		if output := MergeOverlappingIntervals(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
