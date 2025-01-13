package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestPalindromeList(t *testing.T) {

	tests := []struct {
		input  *utils.ListNode
		output int
	}{
		{utils.GenerateLinkedList([]int{1}), 1},
		{utils.GenerateLinkedList([]int{1, 2}), 0},
		{utils.GenerateLinkedList([]int{1, 1}), 1},
		{utils.GenerateLinkedList([]int{1, 2, 1}), 1},
		{utils.GenerateLinkedList([]int{1, 2, 2}), 0},
		{utils.GenerateLinkedList([]int{1, 2, 2, 1}), 1},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := PalindromeList(input)
		if test.output != result {
			t.Errorf("PalindromeList(%v) = %v ; want %v", inputCompare, result, test.output)
		}

	}
}
