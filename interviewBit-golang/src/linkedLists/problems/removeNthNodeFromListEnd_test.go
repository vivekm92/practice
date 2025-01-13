package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input1 *utils.ListNode
		input2 int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), 1, utils.GenerateLinkedList([]int{1, 2, 3, 4})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), 2, utils.GenerateLinkedList([]int{1, 2, 3, 5})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), 6, utils.GenerateLinkedList([]int{2, 3, 4, 5})},
	}

	for _, test := range tests {
		input1, inputCompare1 := test.input1.Copy(), test.input1.Copy()
		result := RemoveNthFromEnd(input1, test.input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("RemoveNthFromEnd(%v, %v) = %v ; want %v", inputCompare1, test.input2, result, test.output)
		}
	}
}
