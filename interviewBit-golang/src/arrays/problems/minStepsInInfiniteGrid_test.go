package arrayProblems

import (
	"fmt"
	"testing"
)

func TestMinStepsInInfiniteGrid(t *testing.T) {

	tests := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{[]int{0, 1, 1}, []int{0, 1, 2}, 2},
		{[]int{1, -2, 4}, []int{0, -2, 5}, 10},
	}

	for idx, test := range tests {
		result := minStepsInInfiniteGrid(test.A, test.B)
		if result != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, result)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, result, test.Expected)
		}
	}
}
