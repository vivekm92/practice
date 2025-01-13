package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestDetectCycle(t *testing.T) {

	tests := []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedListWithCycle([]int{1, 2, 3, 4, 5}, -1), utils.GenerateLinkedList([]int{})},
		{utils.GenerateLinkedListWithCycle([]int{1, 2, 3, 4, 5}, 0), utils.GenerateLinkedListWithCycle([]int{1, 2, 3, 4, 5}, 0)},
		{utils.GenerateLinkedListWithCycle([]int{1, 2, 3, 4, 5}, 1), utils.GenerateLinkedListWithCycle([]int{2, 3, 4, 5}, 0)},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := DetectCycle(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("DetectCycle(%v) = %v ; want %v", inputCompare, result, test.output)
		}
	}
}
