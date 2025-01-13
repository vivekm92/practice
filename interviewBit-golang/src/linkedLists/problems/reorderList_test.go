package linkedList

import (
	"fmt"
	"interviewBit/src/utils"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 5, 2, 4, 3})},
		{utils.GenerateLinkedList([]int{1}), utils.GenerateLinkedList([]int{1})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6, 7}), utils.GenerateLinkedList([]int{1, 7, 2, 6, 3, 5, 4})},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := ReorderList(input)
		fmt.Println(result)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReorderList(%v) = %v; want %v", inputCompare, result, test.output)
		}
	}
}
