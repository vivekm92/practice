package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	tests := []struct {
		input1 *utils.ListNode
		input2 *utils.ListNode
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3}), utils.GenerateLinkedList([]int{4, 5, 6}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6})},
		{utils.GenerateLinkedList([]int{1, 2, 3}), utils.GenerateLinkedList([]int{}), utils.GenerateLinkedList([]int{1, 2, 3})},
		{utils.GenerateLinkedList([]int{}), utils.GenerateLinkedList([]int{4, 5, 6}), utils.GenerateLinkedList([]int{4, 5, 6})},
		{utils.GenerateLinkedList([]int{1, 3, 5}), utils.GenerateLinkedList([]int{2, 4, 6}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6})},
	}

	for _, test := range tests {
		input1, inputCompare1 := test.input1.Copy(), test.input1.Copy()
		input2, inputCompare2 := test.input2.Copy(), test.input2.Copy()
		result := MergeTwoLists(input1, input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("MergeTwoLists(%v, %v) = %v; want %v", inputCompare1.String(), inputCompare2.String(), result.String(), test.output.String())
		}
	}
}
