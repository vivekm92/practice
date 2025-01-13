package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	var tests = []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{3, 2, 1, 4, 0}), utils.GenerateLinkedList([]int{0, 1, 2, 3, 4})},
		{utils.GenerateLinkedList([]int{5, 4, 3, 2, 1}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})},
	}

	for _, test := range tests {
		input, inputCompare := test.input.Copy(), test.input.Copy()
		result := InsertionSortList(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("InsertionSortList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}

}
