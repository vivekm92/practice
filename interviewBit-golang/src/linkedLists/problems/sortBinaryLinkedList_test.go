package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestSortBinaryLinkedList(t *testing.T) {

	var tests = []struct {
		input  *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{0, 0, 1, 1, 0, 1, 1, 1}), utils.GenerateLinkedList([]int{0, 0, 0, 1, 1, 1, 1, 1})},
		{utils.GenerateLinkedList([]int{0, 1, 1}), utils.GenerateLinkedList([]int{0, 1, 1})},
	}

	for _, test := range tests {
		input := test.input.Copy()
		inputCompare := test.input.Copy()
		result := SortBinaryLinkedList(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("SortBinaryLinkedList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}

	for _, test := range tests {
		input := test.input.Copy()
		inputCompare := test.input.Copy()
		result := SortBinaryLinkedList1(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("SortBinaryLinkedList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}

	for _, test := range tests {
		input := test.input.Copy()
		inputCompare := test.input.Copy()
		result := SortBinaryLinkedList2(input)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("SortBinaryLinkedList(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}
}
