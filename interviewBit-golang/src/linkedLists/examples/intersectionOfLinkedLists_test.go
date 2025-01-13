package linkedListExamples

import (
	"interviewBit/src/utils"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	var tests = []struct {
		input1 *utils.ListNode
		input2 *utils.ListNode
		pos1   int
		pos2   int
		output *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), 1, 2, utils.GenerateLinkedList([]int{3, 4, 5})},
		{utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), utils.GenerateLinkedList([]int{1, 2, 3, 4, 5}), 0, 1, utils.GenerateLinkedList([]int{2, 3, 4, 5})},
	}

	for _, test := range tests {
		input1 := test.input1.Copy()
		input2 := test.input2.Copy()
		input1, input2 = utils.MergeLinkedList(input1, input2, test.pos1, test.pos2)
		inputCompare1, inputCompare2 := input1.Copy(), input2.Copy()
		result := GetIntersectionNode(input1, input2)
		if !utils.CompareLinkedLists(result, test.output) {
			t.Errorf("GetIntersectionNode(%v, %v) = %v; want %v", inputCompare1.String(), inputCompare2.String(), result.String(), test.output.String())
		}
	}
}
