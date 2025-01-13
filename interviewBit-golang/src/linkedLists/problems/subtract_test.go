package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestSubtract(t *testing.T) {
	tests := []struct {
		A        *utils.ListNode
		Expected *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{4, 2, 3, 4, 5})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{3, 1, 3, 4})},
	}

	for _, test := range tests {
		input := test.A.Copy()
		result := Subtract(test.A)
		if !utils.CompareLinkedLists(result, test.Expected) {
			t.Errorf("Subtract(%v) = %v ; want %v", input, result, test.Expected)
		}
	}

}
