package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type waveArrayTestCase struct {
	A        []int
	Expected []int
}

var waveArrayTestCases = []waveArrayTestCase{
	{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 1, 4, 3, 5},
	},
}

func TestWave(t *testing.T) {

	for idx, test := range waveArrayTestCases {
		if output := Wave(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

	for idx, test := range waveArrayTestCases {
		if output := WaveArray(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
