package arrayProblems

import (
	"fmt"
	"testing"
)

type hotelBookingsPossibleTestCase struct {
	A        []int
	B        []int
	C        int
	Expected bool
}

var hotelBookingsPossibleTestCases = []hotelBookingsPossibleTestCase{
	{
		[]int{1, 3, 5},
		[]int{2, 6, 8},
		1,
		false,
	},
	{
		[]int{1, 2, 3},
		[]int{2, 3, 4},
		2,
		true,
	},
	{
		[]int{13, 14, 36, 19, 44, 1, 45, 4, 48, 23, 32, 16, 37, 44, 47, 28, 8, 47, 4, 31, 25, 48, 49, 12, 7, 8},
		[]int{28, 27, 61, 34, 73, 18, 50, 5, 86, 28, 34, 32, 75, 45, 68, 65, 35, 91, 13, 76, 60, 90, 67, 22, 51, 53},
		23,
		true,
	},
}

func TestHotelBookingsPossible(t *testing.T) {

	for idx, test := range hotelBookingsPossibleTestCases {
		if output := HotelBookings(test.A, test.B, test.C); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
