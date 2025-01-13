package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestDeleteDuplicatesII(t *testing.T) {
	tests := []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 2, 2, 3}), utils.GenerateLinkedList([]int{1, 3})},
		{utils.GenerateLinkedList([]int{1, 2, 2, 3, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 4, 5})},
		{utils.GenerateLinkedList([]int{2, 2, 3, 3, 4, 5}), utils.GenerateLinkedList([]int{4, 5})},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := DeleteDuplicatesII(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("DeleteDuplicatesII(%v) = %v ; want %v", inputCompare, result, test.output)
		}
	}
}
