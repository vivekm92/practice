package stacksAndQueues

import (
	"testing"
)

func TestTicketCounter(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, 1},
		{[]int{1, 7, 6, 5}, []int{3, 3, 3, 1}, 1},
		{[]int{1}, []int{10}, 0},
	}

	for _, test := range tests {
		result := TicketCounter(test.A, test.B)
		if result != test.Expected {
			t.Errorf("TicketCounter(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
