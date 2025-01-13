package linkedListExamples

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReverseList(t *testing.T) {

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
		result := ReverseList(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}

	for _, test := range tests {
		input := test.input.Copy()
		inputCompare := test.input.Copy()
		result := ReverseList1(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseList1(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}
}
