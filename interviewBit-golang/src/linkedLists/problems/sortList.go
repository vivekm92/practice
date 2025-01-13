package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/sort-list/
*/

// T(n) : O(n), S(n) : O(n)
func MergeSortedLists(A *utils.ListNode, B *utils.ListNode) *utils.ListNode {

	var dummyHead *utils.ListNode = utils.ListNode_new(-1)
	var curr *utils.ListNode = dummyHead
	for A != nil && B != nil {
		if A.Value < B.Value {
			curr.Next = A
			A = A.Next
		} else {
			curr.Next = B
			B = B.Next
		}
		curr = curr.Next
	}

	if A == nil {
		curr.Next = B
	} else {
		curr.Next = A
	}

	return dummyHead.Next
}

// T(n) : O(nlogn), S(n) : O(logn)
func SortList(A *utils.ListNode) *utils.ListNode {

	if A == nil || A.Next == nil {
		return A
	}

	r, mid := A.Next, A
	for r != nil && r.Next != nil {
		mid = mid.Next
		r = r.Next.Next
	}

	left, right := A, mid.Next
	mid.Next = nil
	ll := SortList(left)
	rl := SortList(right)

	return MergeSortedLists(ll, rl)
}
