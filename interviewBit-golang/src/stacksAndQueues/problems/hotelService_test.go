package stacksAndQueues

import (
	"fmt"
	"reflect"
	"testing"
)

type nearestHotelTestCase struct {
	A        [][]int
	B        [][]int
	Expected []int
}

var nearestHotelTestCases = []nearestHotelTestCase{
	{[][]int{{0, 0}, {1, 0}}, [][]int{{1, 1}, {2, 1}, {1, 2}}, []int{1, 0, 2}},
	{[][]int{{1, 0, 0, 1}}, [][]int{{1, 2}, {1, 3}}, []int{1, 1}},
	// {[][]int{{}}, [][]int{}, []int{}},
}

func TestNearestHotel(t *testing.T) {
	for idx, test := range nearestHotelTestCases {
		if output := NearestHotel(test.A, test.B); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
