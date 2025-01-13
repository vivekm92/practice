package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

func TestReplaceWithFactors(t *testing.T) {
	tests := []struct {
		A        *utils.ListNode
		B        int
		Expected *utils.ListNode
	}{
		{utils.GenerateLinkedList([]int{1, 2, 3}), 2, utils.GenerateLinkedList([]int{0, 2, 2})},
		{utils.GenerateLinkedList([]int{3, 4, 5}), 3, utils.GenerateLinkedList([]int{3, 3, 3})},
	}

	for _, test := range tests {
		result := ReplaceWithFactors(test.A, test.B)
		if !utils.CompareLinkedLists(result, test.Expected) {
			t.Errorf("ReplaceWithFactors(%v, %v) : %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
