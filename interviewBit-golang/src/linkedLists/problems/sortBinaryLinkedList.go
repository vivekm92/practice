package linkedList

import "interviewBit/src/utils"

/*
	Problem : https://www.interviewbit.com/problems/sort-binary-linked-list/
*/

// T(n) : O(n), S(n), O(n)
func SortBinaryLinkedList(A *utils.ListNode) *utils.ListNode {

	if A == nil {
		return nil
	}

	cnt0, cnt1 := 0, 0
	var head *utils.ListNode = utils.ListNode_new(-1)
	var curr *utils.ListNode = head
	for A != nil {
		if A.Value == 0 {
			cnt0++
		} else {
			cnt1++
		}
		A = A.Next
	}

	for cnt0 > 0 {
		curr.Next = utils.ListNode_new(0)
		curr = curr.Next
		cnt0--
	}

	for cnt1 > 0 {
		curr.Next = utils.ListNode_new(1)
		curr = curr.Next
		cnt1--
	}

	return head.Next
}

// T(n): O(n), S(n) : O(1)
func SortBinaryLinkedList1(A *utils.ListNode) *utils.ListNode {

	var h *utils.ListNode = A
	var curr *utils.ListNode = A
	for curr != nil && curr.Value == 0 {
		curr = curr.Next
	}

	if curr == nil {
		return h
	}

	var currNext *utils.ListNode = curr.Next
	for currNext != nil {

		if currNext.Value == 0 {
			currNext.Value = curr.Value
			curr.Value = 0
			curr = curr.Next
			currNext = currNext.Next
		}

		for currNext != nil && currNext.Value == 1 {
			currNext = currNext.Next
		}
	}

	return h
}

// T(n): O(n), S(n) : O(1)
func SortBinaryLinkedList2(A *utils.ListNode) *utils.ListNode {
	var h *utils.ListNode = A
	cnt0, cnt1 := 0, 0
	for h != nil {
		if h.Value == 0 {
			cnt0++
		} else {
			cnt1++
		}
		h = h.Next
	}

	h = A
	for h != nil {
		if cnt0 > 0 {
			cnt0--
			h.Value = 0
		} else {
			h.Value = 1
		}
		h = h.Next
	}

	return A
}
