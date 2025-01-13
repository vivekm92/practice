package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestEvenReverse(t *testing.T) {

	tests := []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4}), utils.GenerateLinkedList([]int{1, 4, 3, 2})},
		{utils.GenerateLinkedList([]int{1, 2, 3}), utils.GenerateLinkedList([]int{1, 2, 3})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6}), utils.GenerateLinkedList([]int{1, 6, 3, 4, 5, 2})},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := EvenReverse(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("EvenReverse(%v) = %v ; want %v", inputCompare, result, test.output)
		}
	}

}
