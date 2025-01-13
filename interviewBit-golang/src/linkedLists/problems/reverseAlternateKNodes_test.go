package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReverseAlternate(t *testing.T) {

	tests := []struct {
		input1 *utils.ListNode
		input2 int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{3, 4, 7, 5, 6, 6, 15, 61, 16}), 3, utils.GenerateLinkedList([]int{7, 4, 3, 5, 6, 6, 16, 61, 15})},
		{utils.GenerateLinkedList([]int{3, 4, 7, 5, 6, 6, 15, 61}), 2, utils.GenerateLinkedList([]int{4, 3, 7, 5, 6, 6, 15, 61})},
	}

	for _, test := range tests {
		input1, inputCompare1 := test.input1.Copy(), test.input1.Copy()
		result := ReverseAlternate(input1, test.input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseAlternate(%v, %v) = %v ; want %v", inputCompare1, test.input2, result, test.output)
		}
	}
}
