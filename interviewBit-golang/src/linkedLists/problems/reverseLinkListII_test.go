package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReverseBetween(t *testing.T) {

	tests := []struct {
		input1 *utils.ListNode
		input2 int
		input3 int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), 1, 2, utils.GenerateLinkedList([]int{2, 1, 3, 4})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), 1, 1, utils.GenerateLinkedList([]int{1, 2, 3, 4})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), 2, 3, utils.GenerateLinkedList([]int{1, 3, 2, 4})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), 3, 4, utils.GenerateLinkedList([]int{1, 2, 4, 3})},
	}

	for _, test := range tests {
		input1, inputCompare1 := test.input1.Copy(), test.input1.Copy()
		result := ReverseBetween(input1, test.input2, test.input3)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseBetween(%v, %v, %v) = %v ; want %v", inputCompare1, test.input2, test.input3, result, test.output)
		}
	}
}
