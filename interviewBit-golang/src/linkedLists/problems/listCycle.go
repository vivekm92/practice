package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/list-cycle/
*/

// T(n) : O(n), S(n) : O(1)
func DetectCycle(A *utils.ListNode) *utils.ListNode {

	var slow, fast *utils.ListNode = A, A
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	if fast.Next == fast {
		return fast
	}

	for A != slow {
		A = A.Next
		slow = slow.Next
	}

	return slow
}
