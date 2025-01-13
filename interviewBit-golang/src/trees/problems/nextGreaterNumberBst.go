package trees

import (
	"interviewBit/src/utils"
)

func NextGreaterNumberBst(A *utils.TreeNode, B int) *utils.TreeNode {

	st := make([]*utils.TreeNode, 0)
	for len(st) > 0 || A != nil {
		for A != nil {
			st = append(st, A)
			A = A.Left
		}

		A = st[len(st)-1]
		st = st[:len(st)-1]

		if A.Value > B {
			return A
		}
		A = A.Right
	}

	return nil
}
