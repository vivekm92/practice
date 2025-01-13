package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReverseList(t *testing.T) {

	tests := []struct {
		input1 *utils.ListNode
		input2 int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6}), 2, utils.GenerateLinkedList([]int{2, 1, 4, 3, 6, 5})},
	}

	for _, test := range tests {
		input1, inputCompare1 := test.input1.Copy(), test.input1.Copy()
		result := ReverseList(input1, test.input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("ReverseList(%v, %v) = %v ; want %v", inputCompare1, test.input2, result, test.output)
		}
	}
}
