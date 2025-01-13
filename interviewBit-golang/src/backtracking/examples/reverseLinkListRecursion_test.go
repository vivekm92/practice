package backtrackingExamples

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	var tests = []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{5, 4, 3, 2, 1})},
		{utils.GenerateLinkedList([]int{1}), utils.GenerateLinkedList([]int{1})},
		{utils.GenerateLinkedList([]int{1, 2}), utils.GenerateLinkedList([]int{2, 1})},
	}

	for _, test := range tests {
		input := test.input.Copy()
		inputCompare := test.input.Copy()
		result := ReverseLinkedList(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseLinkedList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}
}
