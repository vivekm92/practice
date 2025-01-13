package arrayProblems

import (
	"fmt"
	"testing"
)

func TestPickFromBothSides(t *testing.T) {

	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{5, -2, 3, 1, 2}, 3, 8},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 2, 3, 200, -1}, 3, 202},
		{[]int{1, 2, -10, 200, -1}, 3, 200},
	}

	for idx, test := range tests {
		result := pickFromBothSides(test.A, test.B)
		if result != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, result)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, result, test.Expected)
		}
	}

}
