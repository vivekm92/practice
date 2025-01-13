package trees

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/symmetric-binary-tree/
*/

// T(n) : O(n) , S(n) : O(n)
func SymmetricBinaryTree(A *utils.TreeNode) int {
	if solveSymmetricBinaryTree(A, A) {
		return 1
	}
	return 0
}

func solveSymmetricBinaryTree(A *utils.TreeNode, B *utils.TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		return false
	}
	return A.Value == B.Value &&
		solveSymmetricBinaryTree(A.Left, B.Right) && solveSymmetricBinaryTree(A.Right, B.Left)
}
