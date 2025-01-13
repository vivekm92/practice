package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	tests := []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{2, 1, 4, 3})},
		{utils.GenerateLinkedList([]int{1, 2, 7}), utils.GenerateLinkedList([]int{2, 1, 7})},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := SwapPairs(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("SwapPairs(%v) = %v ; want %v", inputCompare, result, test.output)
		}
	}
}
