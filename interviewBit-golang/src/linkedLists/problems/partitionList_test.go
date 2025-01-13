package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestPartition(t *testing.T) {

	tests := []struct {
		input1 *utils.ListNode
		input2 int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 4, 3, 2, 5, 2}), 3, utils.GenerateLinkedList([]int{1, 2, 2, 4, 3, 5})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 1, 3}), 2, utils.GenerateLinkedList([]int{1, 1, 2, 3, 3})},
	}

	for _, test := range tests {
		input1, inputCompare := test.input1.Copy(), test.input1.Copy()
		result := Partition(input1, test.input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("Partition(%v) = %v; want %v", inputCompare.String(), result.String(), test.output.String())
		}
	}

}
